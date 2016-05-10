package network

import (
	"fmt"
	"net"
)

func GetIp() (addlist []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, addr := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println(ipnet.IP.String())
				addlist = append(addlist, ipnet.IP.String())
				/*
				  192.168.1.100
				  192.168.152.1
				  10.10.172.1
				*/
			}

		}
	}
	return
}
