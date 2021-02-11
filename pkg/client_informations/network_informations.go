package client_informations

import (
	"fmt"
	"net"

	gateway "github.com/arthurguyader/go-gateway"
	log "github.com/sirupsen/logrus"
	openprocesses "github.com/smolveau/openprocesses"
)

func getInterfaceInformations() ([]InterfaceInformations, error) {
	var interfaces []InterfaceInformations

	itfs, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("net.Interfaces failed <- %v", err)
	}

	for _, itf := range itfs {

		addrs, err := itf.Addrs()
		if err != nil {
			return nil, fmt.Errorf("itf.Addrs failed <- %v", err)
		}

		var ip net.IP
		var mask net.IPMask
		var maskN int

		for _, addr := range addrs {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				mask = v.Mask
			case *net.IPAddr:
				ip = v.IP
				mask = ip.DefaultMask()
			}

			// check if ipv4
			if ip == nil || ip.To4() == nil {
				continue
			}

			gateways, err := gateway.GetGatewaysByInterface(&itf)
			if err != nil {
				log.Warn(err)
			}

			maskN, _ = mask.Size()
			interfaces = append(interfaces, InterfaceInformations{itf.Name,
				ip.String(), maskN, itf.HardwareAddr.String(), gateways, itf.Flags})

			break
		}

	}
	return interfaces, nil
}

func getPortsAndProcesses() ([]PortsAndProcessesInformations, error) {
	var out []PortsAndProcessesInformations
	tmp, err := openprocesses.GetListeningSockets()
	if err != nil {
		return nil, fmt.Errorf("portsandprocess.GetListeningSockets() failed <- %v", err)
	}
	for i := 0; i < len(tmp); i++ {
		o := PortsAndProcessesInformations{Port: tmp[i].Port, Process: tmp[i].Process}
		out = append(out, o)
	}
	return out, nil
}

func GetNetworkInformations() (*NetworkInformations, error) {
	interfaces, err := getInterfaceInformations()
	if err != nil {
		return nil, fmt.Errorf("getInterfaceInformations failed <- %v", err)
	}
	portsAndProcesses, err := getPortsAndProcesses()
	if err != nil {
		return nil, fmt.Errorf("getPortsAndProcesses failed <- %v", err)
	}

	networkInformations := NetworkInformations{interfaces, portsAndProcesses}

	return &networkInformations, nil
}
