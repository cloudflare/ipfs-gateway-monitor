package ipfs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchUnavailableCid(ctx *context.Context, cfg *config.IpfsConfig) (time.Duration, error) {
	cid := cfg.UnavailableCid
	start := time.Now()
	_, err := http.Get(fmt.Sprintf("https://%s/ipfs/%s", ctx.GatewayHost, cid))
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
