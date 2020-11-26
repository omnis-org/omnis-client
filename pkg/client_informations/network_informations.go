package client_informations

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
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

			gateways, err := getGateways(itf.Name)
			if err != nil {
				log.Warn(err)
			}

			maskN, _ = mask.Size()
			interfaces = append(interfaces, InterfaceInformations{itf.Name,
				ip.String(), maskN, itf.HardwareAddr.String(), gateways, itf.Flags})

		}

	}
	return interfaces, nil
}

func GetNetworkInformations() (*NetworkInformations, error) {
	interfaces, err := getInterfaceInformations()
	if err != nil {
		return nil, fmt.Errorf("getInterfaceInformations failed <- %v", err)
	}

	networkInformations := NetworkInformations{interfaces}

	return &networkInformations, nil
}
