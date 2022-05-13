package edgehost

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// Config controls how we connect to our forward proxy, and where we listen for
// connections from the reverse proxy.
type Config struct {
	// Addr is the address of the forward proxy.
	Addr string
	// DefaultLaddr is the default local address that we should request.
	DefaultLaddr string
	// Config is the TLS client config for our connections to the forward proxy.
	TLSConfig *tls.Config

	// Services maps the public-facing address that others should connect to, to
	// the local address that we should listen on for connections from the
	// reverse proxy.
	Services map[string]string
}

// TLSConfig reads and parses a public/private keypair from disk, along with a
// root CA pool. All files must contain PEM encoded data.
func TLSConfig(certfile, keyfile, cafile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		return nil, err
	}

	rawPool, err := ioutil.ReadFile(cafile)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM(rawPool); !ok {
		return nil, fmt.Errorf("unable to parse file for CA certificates")
	}

	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            pool,
		InsecureSkipVerify: true,
	}, nil
}
