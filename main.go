package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

var hosts map[string]*Host = make(map[string]*Host)
var ticker *time.Ticker = time.NewTicker(time.Second * 5)

func main() {
	newHost("a", "google.com")
	newHost("b", "torshiztimes.ir")
	newHost("c", "yahoo.com")
	newHost("c", "kashmarphoto.ir")
	newHost("c", "twitter.com")
	newHost("c", "salam.ir")
	newHost("c", "soundcloud.com")
	newHost("c", "godoc.org")
	newHost("c", "rightrelevance.com")
	newHost("c", "emadghasemi.ir")

	//go startWorkers()

	http.HandleFunc("/", index)
	file := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", file))
	http.ListenAndServe(":3030", nil)
}
func startWorkers() {
	fmt.Println("StartWorker")
	for range ticker.C {
		var wg sync.WaitGroup
		for _, h := range hosts {
			wg.Add(1)
			go func(_h *Host) {
				_h.Ping(&wg)
			}(h)
		}
		wg.Wait()

		// add pings info to db
		fmt.Println("----------------------------")
	}
}

func newHost(name, hostname string) (*Host, error) {
	h := &Host{Name: name}

	netProto := "ip4:icmp"
	if strings.Index(hostname, ":") != -1 {
		netProto = "ip6:ipv6-icmp"
	}
	_, err := net.ResolveIPAddr(netProto, hostname)
	if err != nil {
		fmt.Printf("Error : %q\n", err)
		return h, err
	}
	h.Ip = hostname
	h.NetProto = netProto
	hosts[hostname] = h

	fmt.Println(hostname, "is Ok, added to queue")

	return h, nil
}
