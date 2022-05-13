package dnslink

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
)

func FetchIpnsDomainAsUrlPath(ctx *context.Context, cfg *config.DnslinkConfig) (time.Duration, error) {
	keyName := "ipns-domain-as-url-path"
	ng := ctx.PickNode()
	defer ng.Close()
	nsh := ng.Node().NewShell()
	if err := nsh.KeyReimport(keyName, cfg.IpnsKey); err != nil {
		return 0, fmt.Errorf("KeyReimport error: %v", err)
	}
	if err := nsh.Publish(keyName, cfg.IpnsCid); err != nil {
		return 0, fmt.Errorf("Publish error: %v", err)
	}

	target := cfg.IpnsDomain
	// TODO: Assert that "target" has DNSLink set to the IPNS content.
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("https://%s/ipns/%s", ctx.GatewayHost, target))
	elapsed := time.Since(start)
	if err != nil {
		return elapsed, err
	}
	if resp.StatusCode != http.StatusOK {
		return elapsed, fmt.Errorf("The status code is not correct: expected: %v got: %v", http.StatusOK, resp.StatusCode)
	}
	return elapsed, err
}
