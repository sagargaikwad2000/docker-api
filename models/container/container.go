package container

type Port struct {
	IP          string `json:"iP"`
	PrivatePort int    `json:"privatePort"`
	PublicPort  int    `json:"publicPort"`
	Type        string `json:"type"`
}

type Label struct {
	OrgOpencontainersImageCreated     string `json:"org.opencontainers.image.created"`
	OrgOpencontainersImageDescription string `json:"org.opencontainers.image.description"`
	OrgOpencontainersImageLicenses    string `json:"org.opencontainers.image.licenses"`
	OrgOpencontainersImageRefName     string `json:"org.opencontainers.image.ref.name"`
	OrgOpencontainersImageRevision    string `json:"org.opencontainers.image.revision"`
	OrgOpencontainersImageSource      string `json:"org.opencontainers.image.source"`
	OrgOpencontainersImageTitle       string `json:"org.opencontainers.image.title"`
	OrgOpencontainersImageURL         string `json:"org.opencontainers.image.url"`
	OrgOpencontainersImageVersion     string `json:"org.opencontainers.image.version"`
	Service                           string `json:"service"`
}

type Networks struct {
	IPAMConfig          interface{} `json:"ipamconfig"`
	Links               interface{} `json:"links"`
	Aliases             interface{} `json:"aliases"`
	MacAddress          string      `json:"macaddress"`
	NetworkID           string      `json:"networkid"`
	EndpointID          string      `json:"endpointid"`
	Gateway             string      `json:"gateway"`
	IPAddress           string      `json:"ipaddress"`
	IPPrefixLen         int         `json:"ipprefixlen"`
	IPv6Gateway         string      `json:"ipv6gateway"`
	GlobalIPv6Address   string      `json:"globalipv6address"`
	GlobalIPv6PrefixLen int         `json:"globalipv6prefixlen"`
	DriverOpts          interface{} `json:"driveropts"`
	DNSNames            interface{} `json:"dnsnames"`
}

type Mount struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Driver      string `json:"driver"`
	Mode        string `json:"mode"`
	RW          bool   `json:"rw"`
	Propagation string `json:"propagation"`
}

type Container struct {
	Id              string            `json:"id"`
	Names           []string          `json:"names"`
	Image           string            `json:"image"`
	ImageID         string            `json:"imageid"`
	Command         string            `json:"command"`
	Created         int               `json:"created"`
	Ports           []Port            `json:"ports"`
	Labels          Label             `json:"labels"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	HostConfig      map[string]string `json:"hostconfig"`
	NetworkSettings Networks          `json:"networksettings"`
	Mounts          []Mount           `json:"mounts"`
}
