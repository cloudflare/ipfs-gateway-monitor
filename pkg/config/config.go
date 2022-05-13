package config

type Config struct {
	AvailableCids []string
	GatewayHost   string
	PinataJwt     string
	Nodes         []string

	Ipfs    IpfsConfig
	Ipns    IpnsConfig
	Dnslink DnslinkConfig
}

type IpfsConfig struct {
	UnavailableCid string
}

type IpnsConfig struct {
	UnavailableName string
}

type DnslinkConfig struct {
	IpfsDomain  string
	IpnsDomain  string
	EmptyDomain string

	IpnsKey string
	IpnsCid string
}
