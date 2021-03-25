package bind

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnConnected(data bool)     // Node Host has Bootstrapped
	OnReady(data bool)         // Node Host has Bootstrapped
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnRemoteStart(data []byte) // User started remote
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ Passes binded Methods to Node ^
func (mn *MobileNode) nodeCallback() dt.NodeCallback {
	return dt.NodeCallback{
		// Direct
		Connected:   mn.call.OnConnected,
		Ready:       mn.call.OnReady,
		Refreshed:   mn.call.OnRefreshed,
		Event:       mn.call.OnEvent,
		RemoteStart: mn.call.OnRemoteStart,
		Responded:   mn.call.OnResponded,
		Progressed:  mn.call.OnProgress,

		// Middleware
		Invited:     mn.invited,
		Received:    mn.received,
		Transmitted: mn.transmitted,
		Queued:      mn.queued,
		Error:       mn.error,

		// User
		GetPeer: mn.user.Peer,
	}
}

// ^ queued Callback, Sends File Invite to Peer, and Notifies Client ^
func (mn *MobileNode) queued(card *md.TransferCard, req *md.InviteRequest) {
	// Retreive Current File
	currFile := mn.user.FS.CurrentFile()
	if currFile == nil {
		log.Println("No Current file")
		return
	}

	// Invite With file
	err := mn.node.InviteFile(card, req, mn.user.Peer(), currFile)
	if err != nil {
		log.Println(err)
		return
	}
}

// ^ invite Callback with data for Lifecycle ^ //
func (mn *MobileNode) invited(data []byte) {
	// Update Status
	mn.status = md.Status_INVITED
	// Callback with Data
	mn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (mn *MobileNode) transmitted(peer *md.Peer) {
	// Update Status
	mn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		// sentry.CaptureException(err)
		log.Println(err)
		return
	}

	// Callback with Data
	mn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (mn *MobileNode) received(card *md.TransferCard) {
	// Update Status
	mn.status = md.Status_AVAILABLE

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		// sentry.CaptureException(err)
		log.Println(err)
		return
	}

	// Callback with Data
	mn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (mn *MobileNode) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println(err)
		return
	}
	// Send Callback
	mn.call.OnError(bytes)
}
