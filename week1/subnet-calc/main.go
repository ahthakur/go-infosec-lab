package main

import (
    "fmt"
    "os"
	"net"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: subnet-calc <CIDR>")
        os.Exit(1)
    }
    cidr := os.Args[1]
    fmt.Println("Parsing:", cidr)
	ip, network, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Invalid CIDR",err)
	}
	fmt.Println("IP:", ip)	
	fmt.Println("Network:", network.IP)
	fmt.Println("Mask:", network.Mask)
	
	broadcast := make([]byte, len(network.Mask))
	for i := range network.IP {
		broadcast[i] = network.IP[i]| ^network.Mask[i]
	}
	fmt.Println("Broadcast ", net.IP(broadcast))
	ones, bits := network.Mask.Size()
	host:= 1 << (bits-ones)
	fmt.Println("Total Hosts ", host)

	firstUsable := make([]byte, len(network.IP))
	copy(firstUsable,network.IP)
	firstUsable[len(firstUsable)-1]++
	fmt.Println("firstUsable IP : ", net.IP(firstUsable))

	lastUsale := make([]byte,len(network.IP))
	copy(lastUsale, broadcast)
	lastUsale[len(lastUsale)-1]--
	fmt.Println("lastUsable IP : ", net.IP(lastUsale))

}