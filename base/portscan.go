package main

import (
	"flag"
	"fmt"
	pinger "github.com/paulstuart/ping"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

/*常用端口号*/
var CommonPort = []string{"21", "22", "23", "25", "53", "80",
	"81", "88", "110", "123", "135", "139", "161", "443",
	"445", "1433", "1251", "3306", "3389", "8080", "8081",
	"8088", "8888"}

/*常用端口号*/
var wg sync.WaitGroup
var timeout = 3 * time.Second
var token chan bool
var simple *bool
var ishelp *bool
var args []string
var size int
var speed *int

func init() {
	speed = flag.Int("s", 10, "-s 指定扫描速度,数值越大扫描速度越快")
	simple = flag.Bool("a", false, "-a 输出全部信息,默认只输出成功信息")
	ishelp = flag.Bool("help", false, "显示帮助")
	flag.Parse()
	token = make(chan bool, *speed)
	help()
	args = flag.Args()
	size = len(args)

}
func info() {
	fmt.Println("扫描速度:", *speed)
	fmt.Println("显示未开放的端口:", *simple)
}
func main() {
	switch size {
	case 1:
		//扫描常见端口号
		fmt.Println("扫描常见端口号")
		info()
		ping(args[0])
		for _, v := range CommonPort {
			wg.Add(1)
			go Scan(args[0] + ":" + v)
		}
	case 2:
		//扫描指定端口
		check(args[1])
		*simple = true
		timeout = 1 * time.Second
		//for {
		scan(args[0] + ":" + args[1])
		//time.Sleep(timeout)
		//}
	case 3:
		//扫描指定端口范围
		info()
		start := args[1]
		end := args[2]
		s := check(start)
		e := check(end)
		if s > e {
			s, e = e, s
		}
		for ; s <= e; s++ {
			wg.Add(1)
			go Scan(args[0] + ":" + strconv.Itoa(s))
		}

	default:
		//参数错误
		fmt.Println("参数错误!  -help查看帮助")
		os.Exit(2)
	}
	wg.Wait()
}
func help() {
	if *ishelp {
		//打印帮助信息
		fmt.Println("\tGoPortScan Version 0.2")
		fmt.Println("port.exe [option] IP或域名 端口号(端口范围)\n")
		fmt.Println("-s speed 指定扫描速度,数值越大扫描速度越快,默认为10")
		fmt.Println("-a allinfo 输出全部信息,默认只输出成功信息")
		os.Exit(0)
	}
}
func Scan(ip string) {

	ctrl(ip)
	wg.Done()
}

func ctrl(ip string) {
	token <- true
	scan(ip)
	<-token

}
func scan(ip string) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	start_time := time.Now()
	conn, err := net.DialTimeout("tcp4", ip, timeout)
	if err != nil {
		if !*simple { //不输出失败信息
			return
		}
		//端口未开放
		fmt.Println(ip + " 端口未开放")
		return
	}
	conn.Close()
	timeused := time.Since(start_time)
	if !Retry(ip) {
		return
	}
	fmt.Printf("%s 开放 用时：%v\n", ip, timeused)
}

func Retry(ip string) bool {
	for i := 0; i < 3; i++ {
		if !retry(ip) {
			return false
		}
	}
	return true
}

func retry(ip string) bool {
	conn, err := net.DialTimeout("tcp4", ip, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func check(port string) int {
	p, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("端口参数错误")
		os.Exit(2)
		return 0
	}
	return p
}
func ping(ip string) {
	ch := make(chan bool, 5)
	fmt.Printf("测试5次ping...")
	for i := 0; i < 5; i++ {
		go func(ch chan bool, ip string) {
			err := pinger.Pinger(ip, 3)
			if err != nil {
				if *simple {
					fmt.Println("\n", err)
				}
				ch <- false
			} else {
				ch <- true
			}

		}(ch, ip)
	}
	for i := 0; i < 5; i++ {
		b := <-ch
		if b {
			fmt.Printf(" OK")
		} else {
			fmt.Printf(" x")
		}
	}
	fmt.Println("\n开始端口扫描...")
}
