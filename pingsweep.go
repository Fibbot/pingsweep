package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"sort"
	"time"

	"github.com/tatsushid/go-fastping"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
func runIP() {
	userIP := os.Args[1]
	userCIDR := os.Args[2]
	ipStr := userIP + "/" + userCIDR

	var ipsUp []string

	ip, ipnet, err := net.ParseCIDR(ipStr)
	if err != nil {
		fmt.Println("ok")
	}

	p := fastping.NewPinger()
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		if len(ip.String()) > 0 {
			p.AddIP(ip.String())
		}
	}

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		//fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		ipsUp = append(ipsUp, addr.String())
	}
	p.OnIdle = func() {
		fmt.Println("scan finished")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hosts Up:")
	fmt.Println("================================")
	sortedIPs := make([]net.IP, 0, len(ipsUp))
	for _, ips := range ipsUp {
		sortedIPs = append(sortedIPs, net.ParseIP(ips))
	}

	sort.Slice(sortedIPs, func(i, j int) bool {
		return bytes.Compare(sortedIPs[i], sortedIPs[j]) < 0
	})
	for _, ips := range sortedIPs {
		fmt.Printf("%s\n", ips)
	}

}

func main() {
	/*
		- ./binary 10.10.10.1 24
		- runs a scan on 10.10.10.1/24
	*/
	fmt.Println("Starting scan")
	runIP()

}
