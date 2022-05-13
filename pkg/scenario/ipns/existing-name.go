package ipns

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchExistingName(ctx *context.Context, cfg *config.IpnsConfig) (time.Duration, error) {
	keyName := "existing-name"

	ownerNg := ctx.PickNode()
	defer ownerNg.Close()
	ownerSh := ownerNg.Node().NewShell()

	caterNg := ctx.PickNode()
	defer caterNg.Close()
	caterSh := caterNg.Node().NewShell()

	key, err := ownerSh.KeyRegen(keyName)
	if err != nil {
		return 0, fmt.Errorf("KeyRegen error: %v", err)
	}
	if err := ownerSh.Publish(keyName, ctx.AvailableCids[0]); err != nil {
		return 0, fmt.Errorf("Publish error: %v", err)
	}

	// Cat the content of the name to make sure that it's publicly available in the IPFS network.
	if _, err := caterSh.Cat("/ipns/" + key.Id); err != nil {
		return 0, err
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
