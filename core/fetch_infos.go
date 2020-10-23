package core

import (
	"context"
	"sync"
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
	o := Output{
		OS:         "",
		HostName:   "",
		Interfaces: interfaces,
		OpenPorts:  nil,
	}
	g.safeData.Add(o)
	return g.safeData.out, nil
}
