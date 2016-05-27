package network

import (
	"fmt"
	"net"
	"os"
)

func Look() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip4", name) //net必须是"ip"、"ip4"或"ip6"。
	if err != nil {
		fmt.Println("Resolvtion error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)

}
