package main

import (
	"bufio"
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
		// go func() {
		// 	reader := bufio.NewReader(os.Stdin)
		// 	line, err := reader.ReadString('\n')
		// 	fmt.Println(line)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		os.Exit(0)
		// 	}
		// 	switch line {
		// 	case "quit":
		// 		conn.Close()
		// 	}
		// }()
		done <- struct{}{}
		// signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)

	// if _, err := io.Copy(conn, os.Stdin); err != nil {
	// 	fmt.Println("ok")
	// 	// fmt.Println("err", conn, os.Stdin)
	// 	// log.Fatal(err)
	// } else {
	// 	fmt.Println("no")
	// }

	conn.Close()
	<-done // wait for background goroutine to finish

}

func mustCopy(dst io.Writer, src io.Reader) {
	reader := bufio.NewReader(src)
	line, _ := reader.ReadString('\n')
	fmt.Println(line)
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("err", dst, src)
		log.Fatal(err)
	}

}
