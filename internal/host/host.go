package host

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	noise "github.com/libp2p/go-libp2p-noise"
	psub "github.com/libp2p/go-libp2p-pubsub"
	quic "github.com/libp2p/go-libp2p-quic-transport"
	tls "github.com/libp2p/go-libp2p-tls"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/logger"
	"google.golang.org/protobuf/proto"
)

// SNRHostStat is the host stat info
type SNRHostStat struct {
	ID        peer.ID
	PublicKey string
	PeerID    string
	MultAddr  string
	Address   string
}

// SNRHost is the host wrapper for the Sonr Network
type SNRHost struct {
	host.Host
	// Properties
	ctx context.Context

	// Libp2p
	disc   *dsc.RoutingDiscovery
	kdht   *dht.IpfsDHT
	pubsub *psub.PubSub
}

type HostListenAddr struct {
	Addr   string
	Family string
}

// MultiAddrStr Returns MultiAddr from IP Address
func (la HostListenAddr) MultiAddrStr(port int) string {
	if strings.Contains(la.Family, "6") {
		return fmt.Sprintf("/ip6/%s/tcp/%d", la.Addr, port)
	}
	return fmt.Sprintf("/ip4/%s/tcp/%d", la.Addr, port)
}

// NewHost creates a new host
func NewHost(ctx context.Context, conn common.Connection, listenAddrs ...HostListenAddr) (*SNRHost, error) {
	// Get Device KeyChain Account Key
	privKey, err := device.KeyChain.GetPrivKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get private key", err)
	}

	// Initialize DHT
	var kdhtRef *dht.IpfsDHT
	listenAddrStrs := getListenAddrStrings(listenAddrs...)

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.ListenAddrStrings(listenAddrStrs...),

		libp2p.Security(tls.ID, tls.New),
		libp2p.Security(noise.ID, noise.New),
		libp2p.Transport(quic.NewTransport),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			25,            // Lowwater
			50,            // HighWater,
			time.Minute*5, // GracePeriod
		)),
		libp2p.DefaultStaticRelays(),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			kdhtRef = kdht
			return kdht, nil
		}),
		libp2p.NATPortMap(),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		return nil, logger.Error("Failed to initialize libp2p host", err)
	}

	// Create Pub Sub
	ps, err := psub.NewGossipSub(ctx, h, psub.WithDiscovery(dsc.NewRoutingDiscovery(kdhtRef)))
	if err != nil {
		return nil, logger.Error("Failed to Create new Gossip Sub", err)
	}

	// Create Host
	hn := &SNRHost{
		ctx:    ctx,
		Host:   h,
		kdht:   kdhtRef,
		pubsub: ps,
	}

	// Bootstrap Host
	err = hn.Bootstrap()
	if err != nil {
		return nil, logger.Error("Failed to bootstrap libp2p Host", err)
	}

	// Check for Wifi/Ethernet for MDNS
	if conn == common.Connection_WIFI || conn == common.Connection_ETHERNET {
		// err = hn.MDNS()
		// if err != nil {
		// 	return nil, logger.Error("Failed to start MDNS Discovery", err)
		// }
	}
	return hn, nil
}

// ** ─── Host Info ────────────────────────────────────────────────────────
// KadDHT Returns the Kademlia DHT Instance
func (hn *SNRHost) KadDHT() (*dht.IpfsDHT, error) {
	if hn.kdht == nil {
		return nil, errors.New("Kadht is not initialized")
	}
	return hn.kdht, nil
}

// Pubsub Returns Host Node MultiAddr
func (hn *SNRHost) Pubsub() *psub.PubSub {
	return hn.pubsub
}

// SendMessage writes a protobuf go data object to a network stream
func (hn *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	s, err := hn.NewStream(hn.ctx, id, p)
	if err != nil {
		return logger.Error("Failed to start stream", err)
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		return logger.Error("Failed to marshal pb", err)
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		return logger.Error("Failed to write message to stream.", err)
	}
	return nil
}

// Stat returns the host stat info
func (hn *SNRHost) Stat() (*SNRHostStat, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetPubKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get public key", err)
	}

	// Marshal Public Key
	buf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Return Host Stat
	return &SNRHostStat{
		ID:        hn.ID(),
		PublicKey: string(buf),
		PeerID:    hn.ID().Pretty(),
		MultAddr:  hn.Addrs()[0].String(),
	}, nil
}
