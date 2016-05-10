package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// func main() {
//  conn, err := net.Dial("tcp", "localhost:8000")
//  if err != nil {
//      log.Fatal(err)
//  }
//  defer conn.Close()
//  mustCopy(os.Stdout, conn)
// }
// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
// 	go mustCopy(os.Stdout, conn) //读取服务端的输出
// 	mustCopy(conn, os.Stdin)     //输入
// 	fmt.Println(conn)
// }

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	//fmt.Fprintf(conn, "client") //输入的前面加了client
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		//fmt.Println("1")
		done <- struct{}{}
		// signal the main goroutine
	}()
	mustCopy(conn, os.Stdin) //一直等
	conn.Close()
	<-done // wait for background goroutine to finish

}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Println("start")
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("err", dst, src)
		log.Fatal(err)
	}
}

/*
监听服务器
模拟 telnet


*/
