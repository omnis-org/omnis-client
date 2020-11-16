package client_informations

import (
	"github.com/shirou/gopsutil/host"
)

func GetSystemInformations() (*SystemInformations, error) {

	infos, err := host.Info()

	if err != nil {
		return nil, err
	}

	operatingSystemInformations := OperatingSystemInformations{infos.OS,
		infos.Platform,
		infos.PlatformFamily,
		infos.PlatformVersion,
		infos.KernelVersion}

	isVirtualized := false
	virtualizationSystem := ""
	if infos.VirtualizationRole == "guest" {
		isVirtualized = true
		virtualizationSystem = infos.VirtualizationSystem
	}

	virtualizationInformations := VirtualizationInformations{isVirtualized, virtualizationSystem}

	systemInfos := SystemInformations{operatingSystemInformations, virtualizationInformations, infos.Hostname, infos.HostID}
	return &systemInfos, nil
}
