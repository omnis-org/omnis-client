// +build !linux

package client_informations

import (
	"errors"
	"runtime"
)

func getGateways(itfName string) ([]string, error) {
	return nil, errors.New("Not implemented for OS : " + runtime.GOOS)
}
