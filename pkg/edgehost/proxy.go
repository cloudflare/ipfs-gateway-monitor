package edgehost

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strings"
	"time"
)

// readAddr reads a length-prefixed TCP address (as a string), parses it, and
// returns the parsed address.
func readAddr(conn net.Conn) (*net.TCPAddr, error) {
	buff := make([]byte, 1)
	if _, err := io.ReadFull(conn, buff); err != nil {
		return nil, err
	}
	buff = make([]byte, buff[0])
	if _, err := io.ReadFull(conn, buff); err != nil {
		return nil, err
	}
	addr, err := net.ResolveTCPAddr("tcp", string(buff))
	if err != nil {
		return nil, err
	} else if addr.IP.IsLoopback() || addr.IP.IsUnspecified() {
		return nil, fmt.Errorf("cowardly refusing to resolve distinguished address: %v", addr)
	}
	return addr, nil
}

// readProxyAddr reads a PROXY protocol header.
func readProxyAddr(conn net.Conn) (*net.TCPAddr, error) {
	hdr, done := "", false
	for i := 0; i < 1024; i++ {
		buff := make([]byte, 1)
		if _, err := io.ReadFull(conn, buff); err != nil {
			return nil, err
		} else if buff[0] == '\r' {
			if _, err := io.ReadFull(conn, buff); err != nil {
				return nil, err
			} else if buff[0] != '\n' {
				return nil, fmt.Errorf("malformed proxy protocol header")
			}
			done = true
			break
		}
		hdr += string(buff)
	}
	if !done {
		return nil, fmt.Errorf("proxy protocol header was too long, or malformed")
	}
	fields := strings.Fields(hdr)
	if len(fields) != 6 {
		return nil, fmt.Errorf("proxy protocol header has wrong number of fields")
	} else if fields[0] != "PROXY" {
		return nil, fmt.Errorf("proxy protocol header doesn't have correct prefix")
	}

	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(fields[2], fields[4]))
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// encodeAddr returns the length-prefixed, serialized TCP address.
func encodeAddr(addr net.Addr) []byte {
	s := addr.String()
	if len(s) >= 256 {
		panic("string is too long to encode")
	}
	out := make([]byte, 1+len(s))
	out[0] = byte(len(s))
	copy(out[1:], []byte(s))
	return out
}

// ProxyDial sends a request to a foward proxy, asking it to connect to `raddr`
// with local address `laddr`. The proxy sends back the local address that was
// actually chosen (for cases like TCP port 0) and then starts proxying
// transparently.
func ProxyDial(ctx context.Context, config *Config, laddr, raddr *net.TCPAddr) (net.Conn, error) {
	if laddr.IP.IsUnspecified() && laddr.Port == 0 {
		laddr = &net.TCPAddr{}
	}
	buff := append(encodeAddr(laddr), encodeAddr(raddr)...)

	addrs := strings.Split(config.Addr, ",")
	addr := addrs[rand.Intn(len(addrs))]

	d := &net.Dialer{
		Timeout:   30 * time.Second,
		DualStack: true,
		Cancel:    ctx.Done(),
	}
	conn, err := tls.DialWithDialer(d, "tcp", addr, config.TLSConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to dial forward proxy: %v", err)
	}
	conn.SetDeadline(time.Now().Add(30 * time.Second))
	defer conn.SetDeadline(time.Time{})

	if _, err := conn.Write(buff); err != nil {
		return nil, fmt.Errorf("failed to write forward proxy header: %v", err)
	}
	pladdr, err := readAddr(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read proxy-local address: %v", err)
	}

	return &proxyConn{
		Conn: conn,

		laddr: pladdr,
		raddr: raddr,
	}, nil
}

// proxyListen starts listening for connections from a reverse proxy at a
// pre-chosen address. New connections are prefixed with the remote address of
// the client, and then become transparent proxies.
func proxyListen(config *Config, service string, laddr *net.TCPAddr) (net.Listener, error) {
	oladdr, ok := config.Services[service]
	if !ok {
		return nil, fmt.Errorf("cannot reverse proxy for unknown service: %v", laddr)
	}
	list, err := net.Listen("tcp", oladdr)
	if err != nil {
		return nil, err
	}

	return &proxyListener{
		Listener: list,

		laddr: laddr,
	}, nil
}

type proxyListener struct {
	net.Listener

	laddr net.Addr
}

func (pl *proxyListener) Accept() (net.Conn, error) {
	conn, err := pl.Listener.Accept()
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.SetDeadline(time.Time{})

	praddr, err := readProxyAddr(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &proxyConn{
		Conn: conn,

		laddr: pl.laddr,
		raddr: praddr,
	}, nil
}

func (pl *proxyListener) Addr() net.Addr { return pl.laddr }

type proxyConn struct {
	net.Conn

	laddr, raddr net.Addr
}

func (pc *proxyConn) LocalAddr() net.Addr  { return pc.laddr }
func (pc *proxyConn) RemoteAddr() net.Addr { return pc.raddr }
