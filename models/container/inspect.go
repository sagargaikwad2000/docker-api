package container

import (
	"time"
)

type InspectContainer struct {
	AppArmorProfile string   `json:"apparmorprofile"`
	Args            []string `json:"args"`
	Config          struct {
		AttachStderr bool     `json:"attachstderr"`
		AttachStdin  bool     `json:"attachstdin"`
		AttachStdout bool     `json:"attachstdout"`
		Cmd          []string `json:"cmd"`
		Domainname   string   `json:"domainname"`
		Env          []string `json:"env"`
		Healthcheck  struct {
			Test []string `json:"test"`
		} `json:"healthcheck"`
		Hostname        string              `json:"hostname"`
		Image           string              `json:"image"`
		Labels          map[string]string   `json:"labels"`
		MacAddress      string              `json:"macaddress"`
		NetworkDisabled bool                `json:"networkdisabled"`
		OpenStdin       bool                `json:"openstdin"`
		StdinOnce       bool                `json:"stdinonce"`
		Tty             bool                `json:"tty"`
		User            string              `json:"user"`
		Volumes         map[string]struct{} `json:"volumes"`
		WorkingDir      string              `json:"workingdir"`
		StopSignal      string              `json:"stopsignal"`
		StopTimeout     int                 `json:"stoptimeout"`
	} `json:"config"`
	Created    time.Time `json:"created"`
	Driver     string    `json:"driver"`
	ExecIDs    []string  `json:"execids"`
	HostConfig struct {
		MaximumIOps          int        `json:"maximumiops"`
		MaximumIOBps         int        `json:"maximumiobps"`
		BlkioWeight          int        `json:"blkioweight"`
		BlkioWeightDevice    []struct{} `json:"blkioweightdevice"`
		BlkioDeviceReadBps   []struct{} `json:"blkiodevicereadbps"`
		BlkioDeviceWriteBps  []struct{} `json:"blkiodevicewritebps"`
		BlkioDeviceReadIOps  []struct{} `json:"blkiodevicereadiops"`
		BlkioDeviceWriteIOps []struct{} `json:"blkiodevicewriteiops"`
		ContainerIDFile      string     `json:"containeridfile"`
		CpusetCpus           string     `json:"cpusetcpus"`
		CpusetMems           string     `json:"cpusetmems"`
		CpuPercent           int        `json:"cpupercent"`
		CpuShares            int        `json:"cpushares"`
		CpuPeriod            int        `json:"cpuperiod"`
		CpuRealtimePeriod    int        `json:"cpurealtimeperiod"`
		CpuRealtimeRuntime   int        `json:"cpurealtimeruntime"`
		Devices              []struct{} `json:"devices"`
		DeviceRequests       []struct {
			Driver       string            `json:"driver"`
			Count        int               `json:"count"`
			DeviceIDs    []string          `json:"deviceids"`
			Capabilities [][]string        `json:"capabilities"`
			Options      map[string]string `json:"options"`
		} `json:"devicerequests"`
		IpcMode           string `json:"ipcmode"`
		Memory            int64  `json:"memory"`
		MemorySwap        int64  `json:"memoryswap"`
		MemoryReservation int64  `json:"memoryreservation"`
		OomKillDisable    bool   `json:"oomkilldisable"`
		OomScoreAdj       int    `json:"oomscoreadj"`
		NetworkMode       string `json:"networkmode"`
		PidMode           string `json:"pidmode"`
		PortBindings      map[string][]struct {
			HostIP   string `json:"hostip"`
			HostPort string `json:"hostport"`
		} `json:"portbindings"`
		Privileged      bool `json:"privileged"`
		ReadonlyRootfs  bool `json:"readonlyrootfs"`
		PublishAllPorts bool `json:"publishallports"`
		RestartPolicy   struct {
			Name              string `json:"name"`
			MaximumRetryCount int    `json:"maximumretrycount"`
		} `json:"restartpolicy"`
		LogConfig struct {
			Type string `json:"type"`
		} `json:"logconfig"`
		Sysctls      map[string]string `json:"sysctls"`
		Ulimits      []struct{}        `json:"ulimits"`
		VolumeDriver string            `json:"volumedriver"`
		ShmSize      int64             `json:"shmsize"`
	} `json:"hostconfig"`
	HostnamePath    string `json:"hostnamepath"`
	HostsPath       string `json:"hostspath"`
	LogPath         string `json:"logpath"`
	Id              string `json:"id"`
	Image           string `json:"image"`
	MountLabel      string `json:"mountlabel"`
	Name            string `json:"name"`
	NetworkSettings struct {
		Bridge                 string `json:"bridge"`
		SandboxID              string `json:"sandboxid"`
		HairpinMode            bool   `json:"hairpinmode"`
		LinkLocalIPv6Address   string `json:"linklocalipv6address"`
		LinkLocalIPv6PrefixLen int    `json:"linklocalipv6prefixlen"`
		SandboxKey             string `json:"sandboxkey"`
		EndpointID             string `json:"endpointid"`
		Gateway                string `json:"gateway"`
		GlobalIPv6Address      string `json:"globalipv6address"`
		GlobalIPv6PrefixLen    int    `json:"globalipv6prefixlen"`
		IPAddress              string `json:"ipaddress"`
		IPPrefixLen            int    `json:"ipprefixlen"`
		IPv6Gateway            string `json:"ipv6gateway"`
		MacAddress             string `json:"macaddress"`
		Networks               map[string]struct {
			NetworkID           string `json:"networkid"`
			EndpointID          string `json:"endpointid"`
			Gateway             string `json:"gateway"`
			IPAddress           string `json:"ipaddress"`
			IPPrefixLen         int    `json:"ipprefixlen"`
			IPv6Gateway         string `json:"ipv6gateway"`
			GlobalIPv6Address   string `json:"globalipv6address"`
			GlobalIPv6PrefixLen int    `json:"globalipv6prefixlen"`
			MacAddress          string `json:"macaddress"`
		} `json:"networks"`
	} `json:"networksettings"`
	Path           string `json:"path"`
	ProcessLabel   string `json:"processlabel"`
	ResolvConfPath string `json:"resolvconfpath"`
	RestartCount   int    `json:"restartcount"`
	State          struct {
		Error      string    `json:"error"`
		ExitCode   int       `json:"exitcode"`
		FinishedAt time.Time `json:"finishedat"`
		Health     struct {
			Status        string `json:"status"`
			FailingStreak int    `json:"failingstreak"`
			Log           []struct {
				Start    time.Time `json:"start"`
				End      time.Time `json:"end"`
				ExitCode int       `json:"exitcode"`
				Output   string    `json:"output"`
			} `json:"log"`
		} `json:"health"`
		OOMKilled  bool      `json:"oomkilled"`
		Dead       bool      `json:"dead"`
		Paused     bool      `json:"paused"`
		Pid        int       `json:"pid"`
		Restarting bool      `json:"restarting"`
		Running    bool      `json:"running"`
		StartedAt  time.Time `json:"startedat"`
		Status     string    `json:"status"`
	} `json:"state"`
	Mounts []struct {
		Name        string `json:"name"`
		Source      string `json:"source"`
		Destination string `json:"destination"`
		Driver      string `json:"driver"`
		Mode        string `json:"mode"`
		RW          bool   `json:"rw"`
		Propagation string `json:"propagation"`
	} `json:"mounts"`
}
