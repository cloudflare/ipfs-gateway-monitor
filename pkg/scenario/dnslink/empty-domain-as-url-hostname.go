package dnslink

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchEmptyDomainAsUrlHostname(ctx *context.Context, cfg *config.DnslinkConfig) (time.Duration, error) {
	target := cfg.EmptyDomain
	// TODO: Assert that "target" CNAME's to the gateway but has no DNSLink.
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("https://%s", target))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	// The only expected error is 404. Otherwise, emit the error.
	if resp.StatusCode != http.StatusNotFound {
		return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusNotFound, resp.StatusCode)
	}
	return elapsed, err
}
