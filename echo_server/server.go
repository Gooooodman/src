// echo_server project main.go
package main

import (
	//"io"
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		connFrom := conn.RemoteAddr().String()
		fmt.Println(connFrom)
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
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	//fmt.Println(input.Scan())
	for input.Scan() {
		fmt.Printf(input.Text())
		// if input.Text() == "quit" {
		// 	c.Close()
		// }
		go echo(c, input.Text(), 1*time.Second)
		//这里加入了go 表示不等待上一层的输出
	}
	// NOTE: ignoring potential errors from input.Err()
	//c.Close()
}
