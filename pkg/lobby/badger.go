package lobby

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/sonr-io/core/pkg/proto"
	"google.golang.org/protobuf/proto"
)

// GetPeer returns ONE Peer in Datastore
func (lob *Lobby) GetPeer(queryID string) (pb.PeerInfo, error) {
	// Initialize Object
	peer := pb.PeerInfo{}

	// ** Create Transaction ** //
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// Set Transaction Query
		item, err := txn.Get([]byte(queryID))

		// @ Find Item
		err = item.Value(func(val []byte) error {
			err := proto.Unmarshal(val, &peer)
			if err != nil {
				log.Fatal("unmarshaling error: ", err)
			}
			return nil
		})

		// Check for Error
		if err != nil {
			return err
		}
		return nil
	})

	// Check for Error
	if err != nil {
		fmt.Println("Search Error ", err)
	}
	return peer, nil
}

// GetPeer returns ONE Peer in Datastore
func (lob *Lobby) GetPeerID(idStr string) (peer.ID, error) {
	// Get Lobby PeerID Slice
	lobbyPeers := lob.ps.ListPeers(lob.Code)

	// Get Pub/Sub Topic Peers and Iterate
	for _, id := range lobbyPeers {
		// If Found
		if id.String() == idStr {
			return id, nil
		}
	}
	// Create New Error
	err := errors.New("Peer ID for given query not found")
	return "", err
}

// GetAllPeers returns ALL Peers in Datastore
func (lob *Lobby) GetAllPeers() string {
	// ** Initialize Variables ** //
	var peerSlice []*pb.PeerInfo

	// ** Open Data Store Read Transaction ** //
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// @ Create Iterator
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		// @ Iterate over bucket
		for it.Rewind(); it.Valid(); it.Next() {
			// Get Item and Key
			item := it.Item()

			// Get Item Value
			err := item.Value(func(data []byte) error {
				// Convert Value to String Add to Slice
				peer := pb.PeerInfo{}
				err := proto.Unmarshal(data, &peer)
				if err != nil {
					log.Fatal("unmarshaling error: ", err)
				} else {
					// Add Item value to Slice
					peerSlice = append(peerSlice, &peer)
				}
				return nil
			})

			// Check error
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Check for Error
	if err != nil {
		fmt.Println("Transaction Error ", err)
	}

	// ** Convert slice to bytes ** //
	bytes, err := json.Marshal(peerSlice)
	if err != nil {
		println("Error converting peers to json ", err)
	}

	// Return as string
	return string(bytes)
}

// ^ Checks for Peer in Pub/Sub Topic ^ //
func (lob *Lobby) isPeerInLobby(queryID string) bool {
	// Get Lobby PeerID Slice
	lobbyPeers := lob.ps.ListPeers(lob.Code)

	// Get Pub/Sub Topic Peers and Iterate
	for _, id := range lobbyPeers {
		// If Found
		if id.String() == queryID {
			return true
		}
	}
	// If Not Found
	return false
}

// ^ removePeer deletes a peer from the circle ^
func (lob *Lobby) removePeer(peer *pb.PeerInfo) {
	// Delete peer from datastore
	key := []byte(peer.GetId())
	err := lob.peerDB.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})

	// Check for Error
	if err != nil {
		fmt.Println(err)
	}

	// Send Callback with updated peers
	lob.callback.OnRefreshed(lob.GetAllPeers())
	println("")
}

// ^ updatePeer changes peer values in circle ^
func (lob *Lobby) updatePeer(peer *pb.PeerInfo) {
	// Create Key/Value as Bytes
	key := []byte(peer.GetId())

	// Convert Request to Proto Binary
	value, err := proto.Marshal(peer)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Update peer in DataStore
	err = lob.peerDB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, value)
		err := txn.SetEntry(e)
		return err
	})

	// Check Error
	if err != nil {
		fmt.Println("Error Updating Peer in Badger", err)
	}

	// Send Callback with updated peers
	lob.callback.OnRefreshed(lob.GetAllPeers())
}
