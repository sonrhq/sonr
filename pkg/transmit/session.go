package transmit

import (
	"io"
	"sync"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"
)

// Session is a single entry in the Transmit queue.
type Session struct {
	direction   common.Direction
	from        *common.Peer
	to          *common.Peer
	payload     *common.Payload
	compChan    chan itemResult
	lastUpdated int64
}

// MapItems performs PayloadItemFunc on each item in the Payload.
func (s Session) Items() []*common.Payload_Item {
	return s.payload.GetItems()
}

// Count returns the number of items in the payload.
func (s Session) Count() int {
	return len(s.Items())
}

// Event returns the CompleteEvent for the given session.
func (s Session) Event(success map[int32]bool) *api.CompleteEvent {
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.payload,
		CreatedAt:  s.payload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Success:    success,
	}
}

// IsIncoming returns true if the session is incoming.
func (s Session) IsIncoming() bool {
	return s.direction == common.Direction_INCOMING
}

// IsOutgoing returns true if the session is outgoing.
func (s Session) IsOutgoing() bool {
	return s.direction == common.Direction_OUTGOING
}

// HandleComplete handles the completion of a session item.
func (s Session) HandleComplete(n api.NodeImpl, wg *sync.WaitGroup) {
	success := make(map[int32]bool)
	for {
		select {
		case result := <-s.compChan:
			// Update Success
			success[int32(result.index)] = result.success

			// Complete Wait Group
			logger.Debug("Received Item Result", golog.Fields{"success": result.success})
			wg.Done()

			// Replace Incoming Item
			if result.IsIncoming() {
				s.payload.Items[result.index] = result.item
				s.lastUpdated = int64(time.Now().Unix())
			}

			// Check if Complete
			if result.IsAllCompleted(s.Count()) {
				n.OnComplete(s.Event(success))
				return
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Reader Interface
// -----------------------------------------------------------------------------

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) {
	// Initialize Params
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup

	// Write All Files
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)
		go s.HandleComplete(n, &wg)

		// Configure Reader
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			reader: rs,
		}

		// Create Reader
		handleItemRead(config, s.compChan)
		wg.Wait()
	}

	// Wait for all items to be written
	stream.Close()
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(config itemConfig, compChan chan itemResult) {
	// Create New Writer
	ir := &itemReader{}
	config.ApplyReader(ir)

	// Define Finish Function and Start Channels
	callFinishFunc := func(r bool) {
		ir.doneChan <- r
		compChan <- ir.toResult(r)
	}
	go ir.handleProgress()

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := config.reader.ReadMsg()
		if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			callFinishFunc(false)
			break
		} else {
			// Write Chunk to File
			if err := ir.WriteChunk(buf); err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				callFinishFunc(false)
				break
			}
		}

		// Check if Item is Complete
		if ir.isItemComplete() {
			// Stop Channels
			logger.Debug("Item Read is Complete")

			// Flush Buffer to File
			if err := ir.FlushBuffer(); err != nil {
				logger.Errorf("%s - Failed to Sync File on Read Stream", err)
				callFinishFunc(false)
			}

			// Complete Writing to File
			callFinishFunc(true)
			break
		}
	}
}

// FlushBuffer writes the current buffer to the file.
func (p *itemReader) FlushBuffer() error {
	// Write Buffer to File
	err := p.item.WriteFile(p.buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// handleProgress handles the channels for the ItemReader
func (p *itemReader) handleProgress() {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case <-p.doneChan:
			return
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Writer Interface
// -----------------------------------------------------------------------------

// WriteTo writes the Session to the given stream.
func (s Session) WriteTo(stream network.Stream, n api.NodeImpl) {
	// Initialize Params
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup

	// Create New Writer
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)
		go s.HandleComplete(n, &wg)

		// Configure Writer
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			writer: wc,
		}

		// Create New Writer
		handleItemWrite(config, s.compChan)
		wg.Wait()
	}
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig, compChan chan itemResult) {
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)

	// Define Finish Function and Start Channels
	callFinishFunc := func(r bool) {
		iw.doneChan <- r
		compChan <- iw.toResult(r)
	}
	go iw.handleProgress()

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		iw.doneChan <- false
	}

	// Loop through File
	for {
		c, err := chunker.Next()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				logger.Debug("Chunker has reached end of file.")
				callFinishFunc(true)
				break
			}

			// Unexpected Error
			logger.Errorf("%s - Error reading chunk.", err)
			callFinishFunc(false)
			break
		}

		// Write Chunk to Stream
		if err := iw.WriteChunk(c.Data, c.Length); err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			callFinishFunc(false)
			break
		}

		// Check if Item is Complete
		if iw.isItemComplete() {
			logger.Debug("Item Write is Complete")
			callFinishFunc(true)
			break
		}
	}
}

// handleProgress handles the channels for the ItemReader
func (p *itemWriter) handleProgress() {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case <-p.doneChan:
			return
		}
	}
}
