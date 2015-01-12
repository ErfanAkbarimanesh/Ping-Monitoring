package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"sync"
	"time"
)

type Host struct {
	Ip       string
	Name     string
	NetProto string

	Failure int
	Success int

	Pings []*PingInfo
}

type PingInfo struct {
	Duration time.Duration
	Date     time.Time
	Status   int
	Comment  string
}

func (h *Host) addPingInfo(p *PingInfo) {
	fmt.Println(h.Ip, p.Duration, p.Comment)
	h.Pings = append(h.Pings, p)
}
func (h *Host) Ping(wg *sync.WaitGroup) {

	defer wg.Done()
	p := fastping.NewPinger()

	r, err := net.ResolveIPAddr(h.NetProto, h.Ip)
	if err != nil {
		h.addPingInfo(&PingInfo{0, time.Now(), -1, fmt.Sprintf("Error: %q", err)})
	}

	p.AddIPAddr(r)
	p.OnRecv = func(addr *net.IPAddr, t time.Duration) {
		h.addPingInfo(&PingInfo{t, time.Now(), 1, "Success"})
	}
	err = p.Run()
	if err != nil {
		h.addPingInfo(&PingInfo{0, time.Now(), -1, fmt.Sprintf("Error: %q", err)})
	}
}
