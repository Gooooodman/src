package main

import (
	"io"
	//"bufio"
	//"fmt"
	"log"
	"net"
	//"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
		/*
		   handleConn(conn) 没有go 则只能连一个客户端
		   加入go  则起了很多线程使其支持并发 让每一次handleConn的调用都进入一个独立的goroutine。
		*/
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

/*

保持监听
客户端连接返回当前时间
telnet localhost 8000

*/
// func echo(c net.Conn, shout string, delay time.Duration) {
// 	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", shout)
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t", strings.ToLower(shout))
// }

// func handleConn(c net.Conn) {
// 	input := bufio.NewScanner(c)
// 	for input.Scan() {
// 		echo(c, input.Text(), 1*time.Second)
// 	}
// 	// NOTE: ignoring potential errors from input.Err()
// 	c.Close()
// }
