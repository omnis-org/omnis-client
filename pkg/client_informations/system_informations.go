package client_informations

import (
	"fmt"

	"github.com/omnis-org/omnis-client/internal/version"

	"github.com/shirou/gopsutil/host"
)

/* This program will return the system informations of the client machine
 */

func GetSystemInformations() (*SystemInformations, error) {

	hostInfos, err := host.Info()

	if err != nil {
		return nil, fmt.Errorf("host.Info <- %v", err)
	}

	operatingSystemInformations := OperatingSystemInformations{hostInfos.OS,
		hostInfos.Platform,
		hostInfos.PlatformFamily,
		hostInfos.PlatformVersion,
		hostInfos.KernelVersion}

	isVirtualized := false
	virtualizationSystem := ""
	if hostInfos.VirtualizationRole == "guest" {
		isVirtualized = true
		virtualizationSystem = hostInfos.VirtualizationSystem
	}

	virtualizationInformations := VirtualizationInformations{isVirtualized, virtualizationSystem}

	systemInfos := SystemInformations{&operatingSystemInformations, &virtualizationInformations, hostInfos.Hostname, hostInfos.HostID, version.BuildVersion}
	return &systemInfos, nil
}
