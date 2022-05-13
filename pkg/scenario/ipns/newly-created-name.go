package ipns

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchNewlyCreatedName(ctx *context.Context, cfg *config.IpnsConfig) (time.Duration, error) {
	keyName := "newly-created-name"
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

	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, key.Id))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	if resp.StatusCode != http.StatusOK {
		return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}
	return elapsed, err
}
