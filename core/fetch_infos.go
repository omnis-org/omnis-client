package core

import (
	"context"
	"encoding/json"
	"omnis-client/internal"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type Gatherer struct {
	safeData *SafeData
	Fetcher  IFetcher
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

type IFetcher interface {
	Post(IP string, body []byte) (*internal.HTTPResponse, error)
}

func NewGatherer(fetcher IFetcher) *Gatherer {
	safeData := new(SafeData)
	return &Gatherer{
		safeData: safeData,
		Fetcher:  fetcher,
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

	ps := &PortScanner{
		ip:   "127.0.0.1",
		lock: semaphore.NewWeighted(Ulimit()),
	}
	openports := ps.Start(ctx, 1, 65535, 500*time.Millisecond)

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

func (g Gatherer) SendInfos(serverIP string, o Output) (*internal.HTTPResponse, error) {
	requestBody, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	resp, err := g.Fetcher.Post(serverIP, requestBody)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TODO : Function that create a single file with an int ID of the machine
func CreateIDFile(i int) {

}
