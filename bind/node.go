package sonr

import (
	"fmt"
	"math"
	"sync"

	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	Host           host.Host
	PubSub         *pubsub.PubSub
	Lobby          *lobby.Lobby
	Profile        pb.Profile
	Contact        pb.Contact
	AuthStream     authStreamConn
	TransferStream transferStreamConn
	Callback       *Callback
}

// ^ Sends new proximity/direction update ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(direction float64) bool {
	// ** Initialize ** //
	// Update User Values
	sn.Profile.Direction = math.Round(direction*100) / 100

	// Create Message with Updated Info
	notif := &pb.LobbyMessage{
		Event:  "Update",
		Sender: sn.Profile.HostId,
		Data:   sn.getPeerInfo(),
	}

	// Inform Lobby
	err := sn.Lobby.Publish(notif)
	if err != nil {
		fmt.Println("Error Posting NotifUpdate: ", err)
		return false
	}
	return true
}

// ^ Queue adds a file to Process for Transfer, returns key ^ //
// TODO: Implement an Error Schema with proto
func (sn *Node) Queue(path string) bool {
	// ** Get File Metadata ** //
	meta := file.GetMetadata(path)

	// ** Create Thumbnail ** //
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Run Routine
	var thumb []byte
	go file.GetThumbnail(&wg, meta, thumb)
	wg.Wait()
	print("Thumbnail Size: ")

	// Check Size
	if len(thumb) == 0 {
		return false
	}

	// Store in Profile
	sn.Profile.CurrentFile = &pb.CachedFile{
		Metadata:  meta,
		Thumbnail: thumb,
	}
	return true
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(peerId string) bool {
	// ** Get Required Data **
	peerID, err := sn.Lobby.GetPeerID(peerId)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}

	// ** Get Current File ** //
	cachedFile := sn.Profile.GetCurrentFile()
	if cachedFile == nil {
		fmt.Println(err)
		return false
	}

	// ** Create New Auth Stream **
	err = sn.NewAuthStream(peerID)
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}

	// Create Request Message
	authPbf := &pb.AuthMessage{
		Subject:   pb.AuthMessage_REQUEST,
		Peer:      sn.getPeerInfo(),
		Metadata:  sn.Profile.CurrentFile.GetMetadata(),
		Thumbnail: sn.Profile.CurrentFile.GetThumbnail(),
	}

	// ** Send Invite Message **
	err = sn.AuthStream.Write(authPbf)
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(data []byte) bool {
	// Initialize Event
	response := pb.RespondEvent{}
	err := proto.Unmarshal(data, &response)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
		return false
	}

	// Initialize Response
	authMsg := new(pb.AuthMessage)

	// Check Decision
	if response.Decision == true {
		// Set as Accept
		authMsg = &pb.AuthMessage{
			Subject: pb.AuthMessage_ACCEPT,
			Peer:    sn.getPeerInfo(),
		}
	} else {
		// Set as Decline
		authMsg = &pb.AuthMessage{
			Subject: pb.AuthMessage_DECLINE,
			Peer:    sn.getPeerInfo(),
		}
	}

	// Send Message
	err = sn.AuthStream.Write(authMsg)
	if err != nil {
		return false
	}

	// Succesful
	return true
}
