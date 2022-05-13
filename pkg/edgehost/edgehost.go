// Package edgehost implements IPFS networking over a forward and reverse proxy.
// This lets us use the edge as a dual TCP server+client.
package edgehost

import (
	"context"
	"fmt"
	"net"

	"github.com/ipfs/go-ipfs/core/node/libp2p"

	libp2p2 "github.com/libp2p/go-libp2p"
	p2phost "github.com/libp2p/go-libp2p-core/host"
	peer "github.com/libp2p/go-libp2p-core/peer"
	transport "github.com/libp2p/go-libp2p-core/transport"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	tptu "github.com/libp2p/go-libp2p-transport-upgrader"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	mafmt "github.com/whyrusleeping/mafmt"
)

// New is how this package integrates with the IPFS core. It wraps the default
// host option with our custom transport layers.
func New(config *Config) libp2p.HostOption {
	return func(ctx context.Context, id peer.ID, ps pstore.Peerstore, options ...libp2p2.Option) (p2phost.Host, error) {
		opts := make([]libp2p2.Option, len(options), len(options)+2)
		copy(opts, options)
		opts = append(opts, clearTransports, libp2p2.Transport(NewTransport(config)))
		// TODO: Support websockets

		host, err := libp2p.DefaultHostOption(ctx, id, ps, opts...)
		if err != nil {
			return nil, err
		}

		return host, nil
	}
}

func clearTransports(cfg *libp2p2.Config) error {
	cfg.Transports = nil
	return nil
}

// NOTE: Everything below here is a simplified version of the TCP transport.

type Transport struct {
	config *Config

	Upgrader *tptu.Upgrader
}

var _ transport.Transport = &Transport{}

func NewTransport(config *Config) func(*tptu.Upgrader) *Transport {
	return func(upgrader *tptu.Upgrader) *Transport {
		return &Transport{config: config, Upgrader: upgrader}
	}
}

func (t *Transport) CanDial(addr ma.Multiaddr) bool {
	return mafmt.TCP.Matches(addr)
}

func (t *Transport) maDial(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	la, err := net.ResolveTCPAddr("tcp", t.config.DefaultLaddr)
	if err != nil {
		return nil, err
	}
	ra, err := maddrToTcp(raddr)
	if err != nil {
		return nil, err
	}
	la.Port = 0

	conn, err := ProxyDial(ctx, t.config, la, ra)
	if err != nil {
		return nil, err
	}
	return manet.WrapNetConn(conn)
}

func (t *Transport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	conn, err := t.maDial(ctx, raddr)
	if err != nil {
		return nil, err
	}
	return t.Upgrader.UpgradeOutbound(ctx, t, conn, p)
}

func (t *Transport) maListen(laddr ma.Multiaddr) (manet.Listener, error) {
	la, err := maddrToTcp(laddr)
	if err != nil {
		return nil, err
	}
	list, err := proxyListen(t.config, laddr.String(), la)
	if err != nil {
		return nil, err
	}
	return manet.WrapNetListener(list)
}

func (t *Transport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	list, err := t.maListen(laddr)
	if err != nil {
		return nil, err
	}
	return t.Upgrader.UpgradeListener(t, list), nil
}

func (t *Transport) Protocols() []int {
	return []int{ma.P_TCP}
}

func (t *Transport) Proxy() bool {
	return true
}

func maddrToTcp(addr ma.Multiaddr) (*net.TCPAddr, error) {
	la, err := manet.ToNetAddr(addr)
	if err != nil {
		return nil, err
	}
	latcp, ok := la.(*net.TCPAddr)
	if !ok {
		return nil, fmt.Errorf("not a tcp multiaddr: %s", addr)
	}
	return latcp, nil
}
