package core

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
)

// Source : https://medium.com/@KentGruber/building-a-high-performance-port-scanner-with-golang-9976181ec39d

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

func localInterfaces() ([]Interface, error) {
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

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		log.Panic(err)
	}
	s := strings.TrimSpace(string(out))

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func OpenPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			OpenPort(ip, port, timeout)
		} else {
			return false
		}
		return false
	}

	conn.Close()
	log.Info(port, "open")
	return true
}

func (ps *PortScanner) Start(ctx context.Context, f, l int, timeout time.Duration) []int {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	var openports []int
	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				open := OpenPort(ps.ip, port, timeout)
				if open {
					openports = append(openports, port)
				}
			}
		}(port)
	}
	return openports
}
