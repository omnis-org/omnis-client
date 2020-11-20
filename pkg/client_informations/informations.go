package client_informations

import (
	"fmt"
	"net"

	"github.com/omnis-org/omnis-client/config"
	"github.com/omnis-org/omnis-client/internal/version"
)

type OperatingSystemInformations struct {
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platform_family"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
}

type VirtualizationInformations struct {
	IsVirtualized        bool   `json:"is_virtualized"`
	VirtualizationSystem string `json:"virtualization_system"`
}

type SystemInformations struct {
	OperatingSystem            *OperatingSystemInformations `json:"operating_system"`
	VirtualizationInformations *VirtualizationInformations  `json:"virtualization"`
	Hostname                   string                       `json:"hostname"`
	SerialNumber               string                       `json:"serial_number"`
}

type InterfaceInformations struct {
	Name     string    `json:"name"`
	Ipv4     string    `json:"ipv4"`
	Ipv4Mask int       `json:"ipv4_mask"`
	MAC      string    `json:"mac"`
	Gateways []string  `json:"gateways"`
	Flags    net.Flags `json:"flags"`
}

type NetworkInformations struct {
	Interfaces []InterfaceInformations `json:"interfaces"`
	//OpenPorts  []int                   `json:"open_ports"`
}

type OtherInformations struct {
	Location     string `json:"location"`
	Perimeter    string `json:"perimeter"`
	OmnisVersion string `json:"omnis_version"`
}

type Informations struct {
	SystemInformations  *SystemInformations  `json:"system_informations"`
	NetworkInformations *NetworkInformations `json:"network_informations"`
	OtherInformations   *OtherInformations   `json:"other_informations"`
}

func GetInformations() (*Informations, error) {

	systemInformations, err := GetSystemInformations()

	if err != nil {
		return nil, fmt.Errorf("GetSystemInformations failed <- %v", err)
	}

	networkInformations, err := GetNetworkInformations()

	if err != nil {
		return nil, fmt.Errorf("GetNetworkInformations failed <- %v", err)
	}

	otherInformations := OtherInformations{config.GetConfig().Client.Location, config.GetConfig().Client.Perimeter, version.BuildVersion}

	infos := Informations{systemInformations, networkInformations, &otherInformations}
	return &infos, nil
}
