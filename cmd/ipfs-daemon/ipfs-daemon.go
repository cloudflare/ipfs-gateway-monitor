package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/edgehost"
	oldcmds "github.com/ipfs/go-ipfs/commands"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/corehttp"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/jbenet/goprocess"
	manet "github.com/multiformats/go-multiaddr/net"
)

// Version and BuildTime are filled in during build by the Makefile
var (
	Version   = "N/A"
	BuildTime = "N/A"
)

var (
	proxyAddr  = flag.String("proxy-addr", "", "The address of the forward proxy.")
	proxyLaddr = flag.String("proxy-local-addr", "", "The default local address to request on forward-proxied connections.")
	proxyCert  = flag.String("proxy-cert", "", "File containing the TLS client certificate.")
	proxyKey   = flag.String("proxy-key", "", "File containing the TLS client's private key.")
	proxyCA    = flag.String("proxy-ca", "", "CAs to authenticate the forward proxy with.")

	httpApiHost = flag.String("api-host", "0.0.0.0", "The address the HTTP API will listen on.")
	httpApiPort = flag.String("api-port", "5001", "The port the HTTP API will listen on.")
)

func loadPlugins(repoPath string) (*loader.PluginLoader, error) {
	plugins, err := loader.NewPluginLoader(repoPath)
	if err != nil {
		return nil, fmt.Errorf("error loading plugins: %s", err)
	}

	if err := plugins.Initialize(); err != nil {
		return nil, fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return nil, fmt.Errorf("error initializing plugins: %s", err)
	}
	return plugins, nil
}

func main() {
	flag.Parse()

	repoPath := os.Getenv("IPFS_PATH")
	ipfsPort := os.Getenv("IPFS_PORT")

	// Load all the preloaded plugins. Look at
	// vendor/github.com/ipfs/go-ipfs/plugin/loader/preload_list for all the plugins.
	plugins, err := loadPlugins(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	// Open the repo directory.
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	// The node will also close the repo but there are many places we could
	// fail before we get to that. It can't hurt to close it twice.
	defer repo.Close()

	// Instantiate an IPFS node.
	ncfg := &core.BuildCfg{
		Online: true,
		Repo:   repo,
	}
	// Instantiate a repo. This holds all of the meaningful config for our node.
	if *proxyAddr != "" {
		// If given, configure the node to use our dual forward+reverse proxy on
		// the edge for all networking.
		tlsCfg, err := edgehost.TLSConfig(*proxyCert, *proxyKey, *proxyCA)
		if err != nil {
			log.Fatal(err)
		}
		// go-ipfs requires us to listen on some ports.
		services := make(map[string]string)
		services["/ip4/0.0.0.0/tcp/"+ipfsPort] = "0.0.0.0:" + ipfsPort
		proxyConfig := &edgehost.Config{
			Addr:         *proxyAddr,
			DefaultLaddr: *proxyLaddr,
			TLSConfig:    tlsCfg,
			Services:     services,
		}
		log.Printf("Using the dweb-proxy at %s\n", *proxyAddr)
		ncfg.Host = edgehost.New(proxyConfig)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go awaitSignal(cancel)

	node, err := core.NewNode(ctx, ncfg)
	if err != nil {
		log.Fatal(err)
	}
	node.IsDaemon = true
	defer node.Close()

	// Start "core" plugins. We want to do this *before* starting the HTTP
	// API as the user may be relying on these plugins.
	err = plugins.Start(node)
	if err != nil {
		log.Fatal(err)
	}
	node.Process.AddChild(goprocess.WithTeardown(plugins.Close))

	// Start listening for connections from API clients.
	apiAddr := net.JoinHostPort(*httpApiHost, *httpApiPort)
	apiLis, err := net.Listen("tcp", apiAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HTTP API is listening on %s\n", apiAddr)

	cctx := oldcmds.Context{
		ReqLog:        &oldcmds.ReqLog{},
		ConfigRoot:    repoPath,
		ConstructNode: func() (*core.IpfsNode, error) { return node, nil },
	}

	manetLis, err := manet.WrapNetListener(apiLis)
	if err != nil {
		log.Fatal(err)
	}
	if err := node.Repo.SetAPIAddr(manetLis.Multiaddr()); err != nil {
		log.Fatal(err)
	}
	var opts = []corehttp.ServeOption{
		corehttp.MetricsCollectionOption("api"),
		corehttp.CheckVersionOption(),
		corehttp.CommandsOption(cctx),
		corehttp.WebUIOption,
		corehttp.GatewayOption(false, corehttp.WebUIPaths...),
		corehttp.VersionOption(),
	}

	go func() {
		err := corehttp.Serve(node, apiLis, opts...)
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-node.Process.Closed()
}

// awaitSignal waits for standard termination signals, then exits the process.
func awaitSignal(doneFn func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	log.Printf("Signal received: %v; shutting down...", sig)
	doneFn()

	sig = <-sigs
	log.Fatalf("Signal received: %v; forcing shutdown.", sig)
}
