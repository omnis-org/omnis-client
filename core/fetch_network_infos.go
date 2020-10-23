package core

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

func localAddresses() ([]Interface, error) {
	var output []Interface
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Warn(fmt.Errorf(err.Error()))
		return nil, err
	}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Warn(fmt.Errorf(err.Error()))
			continue
		}
		for _, a := range addrs {

			switch v := a.(type) {
			case *net.IPAddr:
				log.Info(i.Name, v)
				output = append(output, Interface{
					Name:      i.Name,
					IPAddress: v.IP.String(),
					Mask:      v.IP.DefaultMask().String(),
				})

			case *net.IPNet:
				log.Info(i.Name, v)
				output = append(output, Interface{
					Name:      i.Name,
					IPAddress: v.IP.String(),
					Mask:      v.Mask.String(),
				})
			}

		}
	}
	return output, nil
}
