package core

// Output structure to gather machine infos
type Output struct {
	OS         string      `json:"os"`
	HostName   string      `json:"endpoint"`
	Interfaces []Interface `json:"interfaces"`
	OpenPorts  []int       `json:"openports"`
}

type Interface struct {
	Name      string
	IPAddress string
	Mask      string
}
