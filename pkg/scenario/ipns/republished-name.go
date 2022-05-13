package ipns

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchRepublishedName(ctx *context.Context, cfg *config.IpnsConfig) (time.Duration, error) {
	keyName := "republished-name"
	ng := ctx.PickNode()
	defer ng.Close()
	nsh := ng.Node().NewShell()
	key, err := nsh.KeyRegen(keyName)
	if err != nil {
		return 0, fmt.Errorf("KeyRegen error: %v", err)
	}
	if err := nsh.Publish(keyName, ctx.AvailableCids[0]); err != nil {
		return 0, fmt.Errorf("First publish error: %v", err)
	}
	resp, err := http.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, key.Id))
	if err != nil {
		return 0, fmt.Errorf("First request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("First request: The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}

	// Republish the new content.
	if err := nsh.Publish(keyName, ctx.AvailableCids[1]); err != nil {
		return 0, fmt.Errorf("Second publish error: %v", err)
	}

	start := time.Now()
	timeout := 10 * time.Minute
	var elapsed time.Duration
	for {
		resp, err = http.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, key.Id))
		elapsed = time.Since(start)
		if err != nil {
			return elapsed, err
		}
		if resp.StatusCode != http.StatusOK {
			return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
		}
		if elapsed > timeout {
			return elapsed, errors.New("Timeout is reached.")
		}

		// Check the returned content after republishing the content.
		gwContent, err := io.ReadAll(resp.Body)
		if err != nil {
			return elapsed, err
		}
		nodeContent, err := nsh.Cat(ctx.AvailableCids[1])
		if err != nil {
			return elapsed, fmt.Errorf("Cat error: %v", err)
		}
		// If the two contents are equal, the gateway can now see the republish. We can successfully stop here.
		if bytes.Equal(gwContent, nodeContent) {
			break
		}
	}
	return elapsed, err
}
