package ipns

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchUnavailableName(ctx *context.Context, cfg *config.IpnsConfig) (time.Duration, error) {
	name := cfg.UnavailableName
	start := time.Now()
	_, err := http.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, name))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	// TODO: The expected error should be 404, but currently the edge sends 524 response.
	// // The only expected error is 404. Otherwise, emit the error.
	// if resp.StatusCode != 404 {
	// 	return elapsed, errors.New("The status code is not 404")
	// }
	return elapsed, err
}
