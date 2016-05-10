package network

import (
	"fmt"
	"net"
)

func GetLookup(s string) (ips []net.IP) {
	ips, err := net.LookupIP(s)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(ips)
	return
}
