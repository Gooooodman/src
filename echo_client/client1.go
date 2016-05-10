package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
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
//  conn, err := net.Dial("tcp", "localhost:8000")
//  if err != nil {
//      log.Fatal(err)
//  }
//  defer conn.Close()
//  go mustCopy(os.Stdout, conn) //读取服务端的输出
//  mustCopy(conn, os.Stdin)     //输入
//  fmt.Println(conn)
// }

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	//fmt.Fprintf(conn, "client") //输入的前面加了client
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	for {
		fmt.Print("请输入命令:")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		switch line {
		case "quit":
			conn.Close()
		default:
			fmt.Println(line)
			SendRequest(conn, line, done)
		}

	}

	//mustCopy(conn, os.Stdin)

	// if _, err := io.Copy(conn, os.Stdin); err != nil {
	//  fmt.Println("ok")
	//  // fmt.Println("err", conn, os.Stdin)
	//  // log.Fatal(err)
	// } else {
	//  fmt.Println("no")
	// }

	conn.Close()
	<-done // wait for background goroutine to finish

}

// func mustCopy(dst io.Writer, src io.Reader) {
// 	reader := bufio.NewReader(src)
// 	line, _ := reader.ReadString('\n')
// 	fmt.Println(line)
// 	if _, err := io.Copy(dst, src); err != nil {
// 		fmt.Println("err", dst, src)
// 		log.Fatal(err)
// 	}

// }

//发送请求
func SendRequest(conn net.Conn, cmd string, done chan struct{}) {
	// tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7076")
	// checkError(err)
	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// checkError(err)
	go SendData(conn, cmd)
	ReadData(conn, done)
	//conn.Close()
}
func SendData(conn net.Conn, data string) {
	fmt.Println(data)
	buf := []byte(data)
	/*向byte字节里添加结束标记*/
	buf = append(buf, 0)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}

/*读取数据*/
func ReadData(conn net.Conn, done chan struct{}) {
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		//fmt.Println("1")
		// go func() {
		//  reader := bufio.NewReader(os.Stdin)
		//  line, err := reader.ReadString('\n')
		//  fmt.Println(line)
		//  if err != nil {
		//      fmt.Println(err)
		//      os.Exit(0)
		//  }
		//  switch line {
		//  case "quit":
		//      conn.Close()
		//  }
		// }()
		done <- struct{}{}
		// signal the main goroutine
	}()
}
