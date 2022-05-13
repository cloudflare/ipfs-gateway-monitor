package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cloudflare/ipfs-gateway-monitor/pkg/config"
	"github.com/cloudflare/ipfs-gateway-monitor/pkg/measure"
)

// Version and BuildTime are filled in during build by the Makefile
var (
	Version   = "N/A"
	BuildTime = "N/A"
)

type Logger struct {
	w *bufio.Writer
}

func (logger Logger) Write(section string, name string, duration time.Duration, err error) {
	errMsg := "nil"
	if err != nil {
		errMsg = err.Error()
	}
	logger.w.WriteString(fmt.Sprintf("%s,%s,%d,%s\n", section, name, duration.Nanoseconds(), errMsg))
	logger.w.Flush()
}

func main() {
	fmt.Printf("#Version=%s, BuildTime=%s\n", Version, BuildTime)

	cfg := config.Config{}

	flag.StringVar(&cfg.GatewayHost, "gateway-host", "cloudflare-ipfs.com",
		"The hostname of the gateway we will run the scenarios against.")
	flag.StringVar(&cfg.PinataJwt, "pinata-jwt", "",
		"The Pinata JWT to be used to call the Pinata API.")
	flag.StringVar(&cfg.Ipfs.UnavailableCid, "ipfs.unavailable-cid", "",
		"The CID of the content that we assume will be always unavailable.")
	flag.StringVar(&cfg.Ipns.UnavailableName, "ipns.unavailable-name", "",
		"The IPNS name that we assume will be always unavailable.")
	flag.StringVar(&cfg.Dnslink.IpfsDomain, "dnslink.ipfs-domain", "",
		"The domain name that points to the IPFS content.")
	flag.StringVar(&cfg.Dnslink.IpnsDomain, "dnslink.ipns-domain", "",
		"The domain name that points to the IPNS name.")
	flag.StringVar(&cfg.Dnslink.IpnsKey, "dnslink.ipns-key", "",
		"The base64-encoded private key used by the IPNS name which is pointed by the domain specified in dnslink.ipns-domain option.")
	flag.StringVar(&cfg.Dnslink.IpnsCid, "dnslink.ipns-cid", "",
		"The CID which the domain specified in dnslink.ipns-domain option will redirect to.")
	flag.StringVar(&cfg.Dnslink.EmptyDomain, "dnslink.empty-domain", "",
		"The domain name that has CNAME'ed to the gateway but has no DNSLink.")

	flag.Func("node", "The address and port (address:port) of HTTP API of the IPFS node to be used while running the scenarios. If there are multiple nodes available, please specify this option multiple times.",
		func(node string) error {
			cfg.Nodes = append(cfg.Nodes, node)
			return nil
		})

	flag.Func("available-cid", "The CID of the content that we assume will be always available. The CID here must not be the directory CID. If there are multiple CIDs, please specify this option multiple times.",
		func(cid string) error {
			cfg.AvailableCids = append(cfg.AvailableCids, cid)
			return nil
		})
	flag.Parse()

	logger := Logger{bufio.NewWriter(os.Stdout)}
	measure.Run(&cfg, &logger)
}
