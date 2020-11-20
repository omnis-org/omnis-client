package client_informations

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func GetSystemInformations() (*SystemInformations, error) {

	infos, err := host.Info()

	if err != nil {
		return nil, fmt.Errorf("host.Info <- %v", err)
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
