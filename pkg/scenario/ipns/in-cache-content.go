package ipns

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchInCacheContent(ctx *context.Context, cfg *config.IpnsConfig) (time.Duration, error) {
	keyName := "in-cache-content"
	ng := ctx.PickNode()
	defer ng.Close()
	nsh := ng.Node().NewShell()
	key, err := nsh.KeyRegen(keyName)
	if err != nil {
		return 0, fmt.Errorf("KeyRegen error: %v", err)
	}
	if err := nsh.Publish(keyName, ctx.AvailableCids[0]); err != nil {
		return 0, fmt.Errorf("Publish error: %v", err)
	}

	// Create a single TLS connection to be reused in the entire function.
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", ctx.GatewayHost), nil)
	if err != nil {
		return 0, fmt.Errorf("TLS connection error: %v", err)
	}

	// Make an HTTP client, but customize it to use only one connection.
	client := http.Client{
		Transport: &http.Transport{
			// When the client wants to make a connection, always return the
			// same connection.
			DialTLS: func(network, addr string) (net.Conn, error) {
				return conn, nil
			},
		},
	}

	// Make one request first to convince the gateway to put the content into
	// the cache.
	resp, err := client.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, key.Id))
	if err != nil {
		return 0, fmt.Errorf("First request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("First request: The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}
	// Before sending the second request, the response has to be all consumed.
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	start := time.Now()
	resp, err = client.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, key.Id))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	if resp.StatusCode != http.StatusOK {
		return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}
	return elapsed, err
}
