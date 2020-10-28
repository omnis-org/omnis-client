package core

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
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
	interfaces, err := localInterfaces()
	if err != nil {
		return Output{}, err
	}

	systemInfo, err := GetKernelInformation()
	if err != nil {
		return Output{}, err
	}
	log.Info("OS: ", systemInfo.OS)
	log.Info("Core: ", systemInfo.Core)
	log.Info("Platform: ", systemInfo.Platform)
	log.Info("GoOs: ", systemInfo.GoOsVersion)
	log.Info("CPUs: ", systemInfo.CPU)

	ps := &PortScanner{
		ip:   "127.0.0.1",
		lock: semaphore.NewWeighted(Ulimit()),
	}
	openports := ps.Start(ctx, 1, 65535, 500*time.Millisecond)
	log.Info(openports)

	netInfo := NetworkInformation{
		Interfaces: interfaces,
		OpenPorts:  openports,
	}

	o := Output{
		SystemInformation:  systemInfo,
		NetworkInformation: netInfo,
	}
	g.safeData.Add(o)
	return g.safeData.out, nil
}
