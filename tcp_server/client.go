package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	LS   = "LS"
	CD   = "CD"
	PWD  = "PWD"
	QUIT = "QUIT"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7076")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//done := make(chan struct{})
	//go func() {
	for {
		fmt.Print("请输入命令:")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		checkError(err)
		//去掉两端的空格
		line = strings.TrimSpace(line)
		//统一转换成大写字母

		line = strings.ToUpper(line)
		//fmt.Println(line)
		arr := strings.SplitN(line, " ", 2)
		//fmt.Println(arr)
		switch arr[0] {
		case LS:
			SendRequest(conn, LS)
			//SendData(conn, cmd)
		case CD:
			SendRequest(conn, CD+" "+strings.TrimSpace(arr[1]))
		case PWD:
			SendRequest(conn, PWD)
		case QUIT:
			//done <- struct{}{}
			conn.Close()
			fmt.Println("程序退出")
			return
		default:
			fmt.Println("命令错误！")
		}
	}
	//}()
	//<-done
	//conn.Close()
}

//发送请求
func SendRequest(conn net.Conn, cmd string) {
	// tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7076")
	// checkError(err)
	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// checkError(err)
	SendData(conn, cmd)
	fmt.Println(ReadData(conn))
	//conn.Close()
}

/*读取数据*/
func ReadData(conn net.Conn) string {
	var data bytes.Buffer
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		//我们的数据以0做为结束的标记
		if buf[n-1] == 0 {
			//n-1去掉结束标记0
			data.Write(buf[0 : n-1])
			break
		} else {
			data.Write(buf[0:n])
		}
	}
	return string(data.Bytes())
}
func SendData(conn net.Conn, data string) {
	buf := []byte(data)
	/*向byte字节里添加结束标记*/
	buf = append(buf, 0)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
