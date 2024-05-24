package main

/*
import (
	"fmt"
	"net"
	"sync"

	"github.com/go-ping/ping"
)

func pingIP(ip net.IP) bool {
	p, err := ping.NewPinger(ip.String())
	if err != nil {
		return false
	}
	p.Count = 1
	p.Run() // blocks until the ping is done
	stats := p.Statistics()
	return stats.PacketsRecv > 0
}
func scanIPRange(startIP, endIP net.IP, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := startIP[2]; i <= endIP[2]; i++ {
		for j := startIP[3]; j <= endIP[3]; j++ {
			ip := net.IPv4(startIP[0], startIP[1], i, j)
			if pingIP(ip) {
				fmt.Println(ip, "is active")
			}
		}
	}
}
func goMap() {
	startIP := net.ParseIP("192.168.137.1").To4()
	endIP := net.ParseIP("192.168.137.254").To4()
	var wg sync.WaitGroup
	for i := startIP[2]; i <= endIP[2]; i++ {
		wg.Add(1)
		go scanIPRange(startIP, endIP, &wg)
	}
	wg.Wait()
}
*/
