package core

// Output structure to gather machine infos
type Output struct {
	SystemInformation  SystemInformation  `json:"systemInformation"`
	NetworkInformation NetworkInformation `json:"networkInformation"`
}

type SystemInformation struct {
	OS          string `json:"os"`
	HostName    string `json:"hostname"`
	Platform    string `json:"platform"`
	Core        string `json:"core"`
	GoOsVersion string `json:"GOOSVersion"`
	CPU         string `json:"CPU"`
}

type NetworkInformation struct {
	Interfaces []Interface `json:"interfaces"`
	OpenPorts  []int       `json:"openports"`
}

type Interface struct {
	Name      string
	IPAddress string
	Mask      string
}
