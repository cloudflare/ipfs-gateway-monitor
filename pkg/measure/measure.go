package measure

import (
	"sync"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/context"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/measure/dnslink"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/measure/ipfs"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/measure/ipns"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/util"
)

func Run(cfg *config.Config, rw util.ResultWriter) {
	ctx := context.NewContext(cfg)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		dnslink.Run(ctx, &cfg.Dnslink, rw)
	}()
	go func() {
		defer wg.Done()
		ipfs.Run(ctx, &cfg.Ipfs, rw)
	}()
	go func() {
		defer wg.Done()
		ipns.Run(ctx, &cfg.Ipns, rw)
	}()
	wg.Wait()
}
