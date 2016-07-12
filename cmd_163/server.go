package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const (
	LS  = "ls"
	CD  = "cd"
	PWD = "pwd"
)

func main() {
	//在7070端口监听
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8888")
	checkError(err)
	listener, err1 := net.ListenTCP("tcp", tcpAddr)
	checkError(err1)
	//listener, err := net.Listen("tcp", "localhost:7076")
	for {
		//等待客户端的连接
		conn, err2 := listener.Accept()
		connFrom := conn.RemoteAddr().String()
		if err != nil {
			/*通常服务端为一个服务，不会因为错误而退出。出错后，继续
			  等待下一个连接请求*/
			fmt.Println(err2)
			continue
		}
		fmt.Printf("来自客户端 %s 的请求\n", connFrom)
		go ServeClient(conn)
	}
}

func ServeClient(conn net.Conn) {
	defer conn.Close()
	// errc := os.Chdir("F:\\Good\\go\\src\\tcp_server")
	// if errc != nil {
	// 	return
	// }
	for {
		str := ReadData(conn)
		connFrom := conn.RemoteAddr().String()
		if str == "" {
			SendData(conn, "接收数据时出错")
			return
		}
		fmt.Printf("收到%s命令：%s\n", connFrom, str)
		switch str {
		case LS:
			ListDir(conn)
		case PWD:
			Pwd(conn)
		default:
			if str[0:2] == CD {
				Chdir(conn, str[3:])
			} else {
				Command(conn, str)
			}
		}
	}
}
func Chdir(conn net.Conn, s string) {
	err := os.Chdir(s)
	if err != nil {
		SendData(conn, err.Error())
	} else {
		SendData(conn, "OK")
	}
}
func ListDir(conn net.Conn) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		SendData(conn, err.Error())
		return
	}
	var str string
	for i, j := 0, len(files); i < j; i++ {
		f := files[i]
		str += f.Name() + "\t"
		if f.IsDir() {
			str += "dir\r\n"
		} else {
			str += "file\r\n"
		}
	}
	SendData(conn, str)
}

/*读取数据*/
func ReadData(conn net.Conn) string {
	var data bytes.Buffer // 缓冲区
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
	// return string(data.Bytes())
	return data.String() //将未读取部分的字节数据作为字符串返回
}
func SendData(conn net.Conn, data string) {
	buf := []byte(data) //字符流
	/*向byte字节里添加结束标记*/
	buf = append(buf, 0)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}
func Pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		SendData(conn, err.Error())
	} else {
		SendData(conn, s)
	}
}

func Command(conn net.Conn, str string) {
	list := strings.Split(str, " ")
	var out bytes.Buffer
	cmd := exec.Command(list[0], list[1:]...)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Cmd execute error!!info:\n", err.Error())
	}
	SendData(conn, out.String())
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
