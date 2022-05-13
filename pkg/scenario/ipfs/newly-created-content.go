package ipfs

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchNewlyCreatedContent(ctx *context.Context, cfg *config.IpfsConfig) (time.Duration, error) {
	// Create a new content by getting a base64-encoded 32-byte random content.
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return 0, err
	}
	content := base64.StdEncoding.EncodeToString(buf)

	// Add the content to the IPFS network.
	cid, err := ctx.PinataClient.PinFileToIpfs([]byte(content))
	if err != nil {
		return 0, err
	}

	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("https://%s/ipfs/%s", ctx.GatewayHost, cid))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	if resp.StatusCode != http.StatusOK {
		return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}
	return elapsed, err
}
