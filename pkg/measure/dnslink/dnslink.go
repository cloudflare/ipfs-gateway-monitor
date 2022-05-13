package dnslink

import (
	"sync"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/scenario/dnslink"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/util"
)

type Scenario struct {
	name        string
	run         func(ctx *context.Context, cfg *config.DnslinkConfig) (time.Duration, error)
	description string
}

var scenarios = []Scenario{
	{
		name:        "ipfs-domain-as-url-hostname",
		run:         dnslink.FetchIpfsDomainAsUrlHostname,
		description: "Fetch the content with the domain as the URL hostname and the domain points to the immutable IPFS content",
	},
	{
		name:        "ipfs-domain-as-url-path",
		run:         dnslink.FetchIpfsDomainAsUrlPath,
		description: "Fetch the content with the domain as the URL path of the gateway and the domain points to the immutable IPFS content",
	},
	{
		name:        "ipns-domain-as-url-hostname",
		run:         dnslink.FetchIpnsDomainAsUrlHostname,
		description: "Fetch the content with the domain as the URL hostname and the domain points to the mutable IPNS content",
	},
	{
		name:        "ipns-domain-as-url-path",
		run:         dnslink.FetchIpnsDomainAsUrlPath,
		description: "Fetch the content with the domain as the URL path of the gateway and the domain points to the mutable IPNS content",
	},
	{
		name:        "empty-domain-as-url-hostname",
		run:         dnslink.FetchEmptyDomainAsUrlHostname,
		description: "Fetch the content with the domain as the URL hostname and the domain has CNAME record set to the gateway  but has no DNSLink set",
	},
	{
		name:        "empty-domain-as-url-path",
		run:         dnslink.FetchEmptyDomainAsUrlPath,
		description: "Fetch the content with the domain as the URL path of the gateway and the domain has CNAME record set to the gateway  but has no DNSLink set",
	},
}

func Run(ctx *context.Context, cfg *config.DnslinkConfig, rw util.ResultWriter) {
	var wg sync.WaitGroup
	for _, sc := range scenarios {
		wg.Add(1)
		go func(sc Scenario) {
			defer wg.Done()
			duration, err := sc.run(ctx, cfg)
			rw.Write("dnslink", sc.name, duration, err)
		}(sc)
	}
	wg.Wait()
}
