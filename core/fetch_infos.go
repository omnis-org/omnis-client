package core

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Gatherer struct {
	safeData *SafeData
}

type SafeData struct {
	mux sync.Mutex
	out Output
}

func (s *SafeData) Add(d Output) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.out = d
}

type IGatherer interface {
	GatherMachineInfos() (Output, error)
}

func NewGatherer() *Gatherer {
	safeData := new(SafeData)
	return &Gatherer{
		safeData: safeData,
	}
}

func (g Gatherer) GatherMachineInfos(ctx context.Context) (Output, error) {
	interfaces, err := localAddresses()
	if err != nil {
		return Output{}, err
	}

	machineHostname, err := GetMachineHostName()
	if err != nil {
		return Output{}, err
	}
	log.Info(machineHostname)

	osInfo, err := GetKernelInformation()
	if err != nil {
		return Output{}, err
	}

	o := Output{
		OS:          osInfo[0],
		HostName:    machineHostname,
		Core:        osInfo[1],
		Platform:    osInfo[2],
		GoOsVersion: osInfo[3],
		CPU:         osInfo[4],
		Interfaces:  interfaces,
		OpenPorts:   nil,
	}
	g.safeData.Add(o)
	return g.safeData.out, nil
}
