package ipns

import (
	"sync"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/scenario/ipns"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/util"
)

type Scenario struct {
	name        string
	run         func(ctx *context.Context, cfg *config.IpnsConfig) (time.Duration, error)
	description string
}

var scenarios = []Scenario{
	{
		name:        "in-cache-content",
		run:         ipns.FetchInCacheContent,
		description: "Fetch the content already existing in the Gateway cache",
	},
	{
		name:        "newly-created-name",
		run:         ipns.FetchNewlyCreatedName,
		description: "Fetch the content of the name newly created in the IPFS network",
	},
	{
		name:        "republished-name",
		run:         ipns.FetchRepublishedName,
		description: "Fetch the content of the name after republishing the new content on that name",
	},
	{
		name:        "existing-name",
		run:         ipns.FetchExistingName,
		description: "Fetch the content of the name already in the IPFS network but not in the Gateway cache yet",
	},
	{
		name:        "unavailable-name",
		run:         ipns.FetchUnavailableName,
		description: "Fetch the content of the name not existing in the IPFS network",
	},
}

func Run(ctx *context.Context, cfg *config.IpnsConfig, rw util.ResultWriter) {
	var wg sync.WaitGroup
	for _, sc := range scenarios {
		wg.Add(1)
		go func(sc Scenario) {
			defer wg.Done()
			duration, err := sc.run(ctx, cfg)
			rw.Write("ipns", sc.name, duration, err)
		}(sc)
	}
	wg.Wait()
}
