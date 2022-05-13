package ipfs

import (
	"sync"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/scenario/ipfs"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/util"
)

type Scenario struct {
	name        string
	run         func(ctx *context.Context, cfg *config.IpfsConfig) (time.Duration, error)
	description string
}

var scenarios = []Scenario{
	{
		name:        "in-cache-content",
		run:         ipfs.FetchInCacheContent,
		description: "Fetch the content already existing in the Gateway cache",
	},
	{
		name:        "newly-created-content",
		run:         ipfs.FetchNewlyCreatedContent,
		description: "Fetch the content newly created in the IPFS network",
	},
	{
		name:        "unavailable-cid",
		run:         ipfs.FetchUnavailableCid,
		description: "Fetch the content not existing in the IPFS network",
	},
}

func Run(ctx *context.Context, cfg *config.IpfsConfig, rw util.ResultWriter) {
	var wg sync.WaitGroup
	for _, sc := range scenarios {
		wg.Add(1)
		go func(sc Scenario) {
			defer wg.Done()
			duration, err := sc.run(ctx, cfg)
			rw.Write("ipfs", sc.name, duration, err)
		}(sc)
	}
	wg.Wait()
}
