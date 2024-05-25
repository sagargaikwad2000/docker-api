package models

type Port struct {
	IP          string `json:"IP"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
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
	IPAMConfig          interface{} `json:"IPAMConfig"`
	Links               interface{} `json:"Links"`
	Aliases             interface{} `json:"Aliases"`
	MacAddress          string      `json:"MacAddress"`
	NetworkID           string      `json:"NetworkID"`
	EndpointID          string      `json:"EndpointID"`
	Gateway             string      `json:"Gateway"`
	IPAddress           string      `json:"IPAddress"`
	IPPrefixLen         int         `json:"IPPrefixLen"`
	IPv6Gateway         string      `json:"IPv6Gateway"`
	GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
	DriverOpts          interface{} `json:"DriverOpts"`
	DNSNames            interface{} `json:"DNSNames"`
}

type Mount struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Driver      string `json:"Driver"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
	Propagation string `json:"Propagation"`
}

type Container struct {
	Id              string            `json:"Id"`
	Names           []string          `json:"Names"`
	Image           string            `json:"Image"`
	ImageID         string            `json:"ImageID"`
	Command         string            `json:"Command"`
	Created         int               `json:"Created"`
	Ports           []Port            `json:"Ports"`
	Labels          Label             `json:"Labels"`
	State           string            `json:"State"`
	Status          string            `json:"Status"`
	HostConfig      map[string]string `json:"HostConfig"`
	NetworkSettings Networks          `json:"NetworkSettings"`
	Mounts          []Mount           `json:"Mounts"`
}
