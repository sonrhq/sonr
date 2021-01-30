package sonr

import (
	"log"
	"math"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ^ Info returns ALL Peer Data as Bytes^
func (sn *Node) Info() []byte {
	// Convert to bytes to view in plugin
	data, err := proto.Marshal(sn.peer)
	if err != nil {
		log.Println("Error Marshaling Lobby Data ", err)
		return nil
	}
	return data
}

// ^ Link with a QR Code ^ //
func (sn *Node) LinkDevice(peerString string) error {
	// Convert String to Bytes
	peer := md.Peer{}

	// Convert to Peer Protobuf
	err := protojson.Unmarshal([]byte(peerString), &peer)
	if err != nil {
		return err
	}

	// TODO: Save Device to Disk
	return nil
}

// ^ Peer returns Current Peer Info ^
func (sn *Node) Peer() *md.Peer {
	return sn.peer
}

// ^ Updates Current Contact Card ^
func (sn *Node) SetContact(conBytes []byte) {
	newContact := &md.Contact{}
	err := proto.Unmarshal(conBytes, newContact)
	if err != nil {
		log.Println(err)
	}
	sn.contact = newContact
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (sn *Node) Update(direction float64) {
	// ** Initialize ** //
	// Update User Values
	sn.peer.Direction = math.Round(direction*100) / 100

	// Set Antipodal
	if direction > 180 {
		sn.peer.Antipodal = math.Round((direction-180)*100) / 100
	} else {
		sn.peer.Antipodal = math.Round((direction+180)*100) / 100
	}

	// Inform Lobby
	err := sn.lobby.Update()
	if err != nil {
		sn.error(err, "Update")
	}
}

// ^ Process adds generates Preview with Thumbnail ^ //
func (sn *Node) Process(procBytes []byte) {
	// Initialize from Info
	request := &md.ProcessRequest{}
	err := proto.Unmarshal(procBytes, request)
	if err != nil {
		log.Println(err)
	}

	// Create Preview
	safePrev := sf.NewPreview(request, sn.call.OnQueued, sn.error)
	sn.files = append(sn.files, safePrev)
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(decision bool) {
	// @ Check Decision

	// Send Response on PeerConnection
	sn.peerConn.Authorize(decision, sn.contact, sn.peer)
}

// ^ Reset Current Queued File Metadata ^ //
func (sn *Node) ResetFile() {
	// Reset Files Slice
	sn.files = nil
	sn.files = make([]*sf.SafePreview, maxFileBufferSize)
}
