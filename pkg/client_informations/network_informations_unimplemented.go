// +build !linux

package client_informations

import (
	"errors"
	"runtime"
)

// TODO : GATEWAYS GATHERING FOR OTHER OS THAN LINUX
func getGateways(itfName string) ([]string, error) {
	return nil, errors.New("Not implemented for OS : " + runtime.GOOS)
}
