package client_informations

import "github.com/vishvananda/netlink"

func getGateways(itfName string) ([]string, error) {
	link, err := netlink.LinkByName(itfName)

	if err != nil {
		return nil, err
	}

	routes, err := netlink.RouteList(link, netlink.FAMILY_V4)

	if err != nil {
		return nil, err
	}

	var gateways []string

	for _, route := range routes {
		if route.Gw != nil {
			gateways = append(gateways, route.Gw.String())
		}

	}
	return gateways, nil
}
