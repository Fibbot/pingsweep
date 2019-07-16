package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/tatsushid/go-fastping"
)

func findIP(input string) string {
	//https://www.socketloop.com/tutorials/golang-find-ip-address-from-string
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}
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
	fmt.Println("Hosts Up:")
	fmt.Println("================================")
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("================================")
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	fmt.Println("./binary 10.10.10.10 31")
	fmt.Println("This will scan 10.10.10.10/31 and output all IP's in that range that are up")
	fmt.Println("Due to running w/ raw sockets may have to sudo/run as root")
	runIP()

}
