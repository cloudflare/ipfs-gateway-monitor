package context

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	ipfsNode "github.com/cloudflare/ipfs-gateway-monitor/pkg/node"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/pinata"
	"golang.org/x/sync/semaphore"
)

type nodeInfo struct {
	occupied bool
	node     ipfsNode.Node
}

type Context struct {
	AvailableCids []string
	GatewayHost   string
	PinataClient  pinata.Client

	// Mutex used to read and write the node infos.
	mu        sync.Mutex
	nodeInfos []nodeInfo
	// Semaphore used to show the number of unoccupied nodes.
	sem *semaphore.Weighted
}

func NewContext(cfg *config.Config) *Context {
	var infos []nodeInfo
	for _, node := range cfg.Nodes {
		info := nodeInfo{
			occupied: false,
			node:     ipfsNode.Node(node),
		}
		infos = append(infos, info)
	}
	return &Context{
		AvailableCids: cfg.AvailableCids,
		GatewayHost:   cfg.GatewayHost,
		PinataClient:  *pinata.NewClient(cfg.PinataJwt),

		nodeInfos: infos,
		sem:       semaphore.NewWeighted(int64(len(cfg.Nodes))),
	}
}

func (c *Context) PickNode() *NodeGuard {
	c.sem.Acquire(context.Background(), 1)
	c.mu.Lock()
	var ng NodeGuard
	ng.mu = &c.mu
	ng.sem = c.sem
	for idx, _ := range c.nodeInfos {
		if !c.nodeInfos[idx].occupied {
			c.nodeInfos[idx].occupied = true
			ng.info = &c.nodeInfos[idx]
			fmt.Printf("#Occupied IPFS node %d\n", idx)
			break
		}
	}
	c.mu.Unlock()
	return &ng
}

type NodeGuard struct {
	info *nodeInfo
	// A pointer pointing to the mutex in Context.
	mu *sync.Mutex
	// A pointer pointing to the semaphore in Context.
	sem *semaphore.Weighted
}

func (ng *NodeGuard) Node() *ipfsNode.Node {
	return &ng.info.node
}

func (ng *NodeGuard) Close() error {
	ng.mu.Lock()
	ng.info.occupied = false
	ng.mu.Unlock()
	ng.sem.Release(1)
	return nil
}
