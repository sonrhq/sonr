package node

import (
	"container/list"
	"context"
	"errors"
	"strings"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*state.Emitter

	// Host and context
	host *host.SNRHost

	// Properties
	ctx context.Context

	// Node Options
	options nodeOptions

	// Persistent Database
	store *bolt.DB

	// Queue - the transfer queue
	queue *list.List

	// Channels
	// TransferProtocol - decisionEvents
	decisionEvents chan *common.DecisionEvent

	// LobbyProtocol - refreshEvents
	refreshEvents chan *common.RefreshEvent

	// MailboxProtocol - mailEvents
	mailEvents chan *common.MailboxEvent

	// TransferProtocol - inviteEvents
	inviteEvents chan *common.InviteEvent

	// TransferProtocol - progressEvents
	progressEvents chan *common.ProgressEvent

	// TransferProtocol - completeEvents
	completeEvents chan *common.CompleteEvent
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, options ...NodeOption) (*Node, *InitializeResponse, error) {
	// Set Node Options
	opts := defaultNodeOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Initialize Host
	host, err := host.NewHost(ctx, host.WithConnection(opts.connection))
	if err != nil {
		return nil, nil, logger.Error("Failed to initialize host", err)
	}

	// Create Node
	node := &Node{
		Emitter:        state.NewEmitter(2048),
		host:           host,
		ctx:            ctx,
		queue:          list.New(),
		decisionEvents: make(chan *common.DecisionEvent),
		refreshEvents:  make(chan *common.RefreshEvent),
		inviteEvents:   make(chan *common.InviteEvent),
		mailEvents:     make(chan *common.MailboxEvent),
		progressEvents: make(chan *common.ProgressEvent),
		completeEvents: make(chan *common.CompleteEvent),
	}

	// Open Database
	err = node.openStore(ctx, host, node.Emitter)
	if err != nil {
		return nil, nil, logger.Error("Failed to open database", err)
	}

	// Begin Background Tasks
	opts.Apply(ctx, node)
	go node.handleEmitter()
	return node, node.newInitResponse(nil), nil
}

// Peer method returns the peer of the node
func (n *Node) Peer() (*common.Peer, error) {
	// Get Profile
	profile, err := n.GetProfile()
	if err != nil {
		return nil, err
	}

	// Get Public Key
	pubKey, err := device.KeyChain.GetSnrPubKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get Public Key", err)
	}

	// Find PublicKey Buffer
	deviceStat, err := device.Stat()
	if err != nil {
		return nil, logger.Error("Failed to get device Stat", err)
	}

	// Marshal Public Key
	pubBuf, err := pubKey.Buffer()
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Return Peer
	return &common.Peer{
		SName:     strings.ToLower(profile.GetSName()),
		Status:    common.Peer_ONLINE,
		Profile:   profile,
		PublicKey: pubBuf,
		Device: &common.Peer_Device{
			HostName: deviceStat.HostName,
			Os:       deviceStat.Os,
			Id:       deviceStat.Id,
			Arch:     deviceStat.Arch,
		},
	}, nil
}

// GetProfile returns the profile for the user from diskDB
func (n *Node) GetProfile() (*common.Profile, error) {
	// Check if Store is open
	if n.store == nil {
		return nil, logger.Error("Failed to Get Profile", ErrStoreNotCreated)
	}

	var profile common.Profile
	err := n.store.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(USER_BUCKET)

		// Check if bucket exists
		if b == nil {
			return ErrProfileNotCreated
		}

		// Get profile buffer
		buf := b.Get(PROFILE_KEY)
		if buf == nil {
			return nil
		}

		// Unmarshal profile
		err := proto.Unmarshal(buf, &profile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Get Profile
	profile, err := n.GetProfile()
	if err != nil {
		return err
	}

	// Create Transfer
	payload, err := common.NewPayload(profile, paths)
	if err != nil {
		return logger.Error("Failed to Supply Paths", err)
	}

	// Add items to transfer
	n.queue.PushBack(payload)
	return nil
}

// Share a peer to have a transfer
func (n *Node) NewRequest(to *common.Peer) (peer.ID, *transfer.InviteRequest, error) {
	// Fetch Element from Queue
	elem := n.queue.Front()
	if elem != nil {
		// Get Payload
		payload := n.queue.Remove(elem).(*common.Payload)

		// Create New ID for Invite
		id, err := device.KeyChain.CreateUUID()
		if err != nil {
			return "", nil, logger.Error("Failed to create new id for Shared Invite", err)
		}

		// Create new Metadata
		meta, err := device.KeyChain.CreateMetadata(n.host.ID())
		if err != nil {
			return "", nil, logger.Error("Failed to create new metadata for Shared Invite", err)
		}

		// Fetch User Peer
		from, err := n.Peer()
		if err != nil {
			return "", nil, logger.Error("Failed to get Node Peer Object", err)
		}

		// Create Invite Request
		req := &transfer.InviteRequest{
			Payload:  payload,
			Metadata: common.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
			Uuid:     common.SignedUUIDToProto(id),
		}

		// Fetch Peer ID from Public Key
		toId, err := to.PeerID()
		if err != nil {
			return "", nil, logger.Error("Failed to fetch peer id from public key", err)
		}
		return toId, req, nil
	}
	return "", nil, errors.New("No items in Transfer Queue.")
}

// Respond to an invite request
func (n *Node) NewResponse(decs bool, to *common.Peer) (peer.ID, *transfer.InviteResponse, error) {
	// Create new Metadata
	meta, err := device.KeyChain.CreateMetadata(n.host.ID())
	if err != nil {
		return "", nil, logger.Error("Failed to create new metadata for Shared Invite", err)
	}

	// Fetch User Peer
	from, err := n.Peer()
	if err != nil {
		return "", nil, err
	}

	// Create Invite Response
	resp := &transfer.InviteResponse{
		Decision: decs,
		Metadata: common.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.PeerID()
	if err != nil {
		return "", nil, logger.Error("Failed to fetch peer id from public key", err)
	}
	return toId, resp, nil
}

// Stat returns the Node info as StatResponse
func (n *Node) Stat() (*StatResponse, error) {
	// Get Profile
	profile, err := n.GetProfile()
	if err != nil {
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Get Host Stats
	hStat, err := n.host.Stat()
	if err != nil {
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
			SName:   profile.SName,
			Profile: profile,
		}, logger.Error("Failed to get Host Stat", err)
	}

	// Get Device Stat
	dStat, err := device.Stat()
	if err != nil {
		return &StatResponse{
			Success: false,
			Error:   err.Error(),
			SName:   profile.SName,
			Profile: profile,
			Network: &StatResponse_Network{
				PublicKey: hStat.PublicKey,
				PeerID:    hStat.PeerID,
				Multiaddr: hStat.MultAddr,
			},
		}, logger.Error("Failed to get Device Stat", err)
	}

	// Return StatResponse
	return &StatResponse{
		SName:   profile.SName,
		Profile: profile,
		Network: &StatResponse_Network{
			PublicKey: hStat.PublicKey,
			PeerID:    hStat.PeerID,
			Multiaddr: hStat.MultAddr,
		},
		Device: &StatResponse_Device{
			Id:        dStat.Id,
			Name:      dStat.HostName,
			Os:        dStat.Os,
			Arch:      dStat.Arch,
			IsDesktop: dStat.IsDesktop,
			IsMobile:  dStat.IsMobile,
		},
	}, nil
}

// HandleEmitter handles the emitter events.
func (n *Node) handleEmitter() {
	for {
		// // Handle Mail Event
		// n.On(mailbox.Event_MAIL_RECEIVED, func(e *state.Event) {
		// 	event := e.Args[0].(*common.MailboxEvent)
		// 	n.mailEvents <- event
		// })

		// Handle Transfer Invite
		n.On(transfer.Event_INVITED, func(e *state.Event) {
			event := e.Args[0].(*common.InviteEvent)
			n.inviteEvents <- event
		})

		// Handle Transfer Decision
		n.On(transfer.Event_RESPONDED, func(e *state.Event) {
			event := e.Args[0].(*common.DecisionEvent)
			n.decisionEvents <- event
		})

		// Handle Transfer Progress
		n.On(transfer.Event_PROGRESS, func(e *state.Event) {
			event := e.Args[0].(*common.ProgressEvent)
			n.progressEvents <- event
		})

		// Handle Transfer Completed
		n.On(transfer.Event_COMPLETED, func(e *state.Event) {
			event := e.Args[0].(*common.CompleteEvent)
			// Check Direction
			if event.Direction == common.CompleteEvent_INCOMING {
				// Add Sender to Recents
				err := n.AddRecent(event.GetFrom().GetProfile())
				if err != nil {
					logger.Error("Failed to add sender's profile to store.", err)
				}
			} else {
				// Add Receiver to Recents
				err := n.AddRecent(event.GetTo().GetProfile())
				if err != nil {
					logger.Error("Failed to add receiver's profile to store.", err)
				}
			}
			n.completeEvents <- event
		})

		// Handle Lobby Join Events
		n.On(lobby.Event_LIST_REFRESH, func(e *state.Event) {
			refreshEvent := e.Args[0].(*common.RefreshEvent)
			n.refreshEvents <- refreshEvent
		})
	}
}
