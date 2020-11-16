package client_informations

import (
	"net"

	log "github.com/sirupsen/logrus"
)

func getInterfaceInformations() ([]InterfaceInformations, error) {
	var interfaces []InterfaceInformations

	itfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, itf := range itfs {

		addrs, err := itf.Addrs()
		if err != nil {
			return nil, err
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
				log.Error(err)
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
		return nil, err
	}

	networkInformations := NetworkInformations{interfaces}

	return &networkInformations, nil
}
