package main

import (
	//"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//执行ls命令  exec.Command("/bin/ls", "-l", "/")
	//执行who命令 exec.Command("/bin/sh","-c","wh")
	//  cmd := exec.Command("/bin/ls", "-l", "/")
	cmd := exec.Command("python", "C:\\Users\\Administrator\\Desktop\\1.py")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Cmd execute error!!info:\n", err.Error())
	}

	fmt.Println(out.String())
	//fmt.Println(cmd.Stdout.)
}

//使用bufio 不好实现
// func main() {
// 	//执行ls命令  exec.Command("/bin/ls", "-l", "/")
// 	//执行who命令 exec.Command("/bin/sh","-c","wh")
// 	//  cmd := exec.Command("/bin/ls", "-l", "/")
// 	cmd := exec.Command("python", "C:\\Users\\Administrator\\Desktop\\1.py")
// 	//var out bufio.ReadWriter
// 	out := bufio.ReadWriter{&bufio.Reader{}, &bufio.Writer{}}
// 	//out = {cmd.Stdout,nil}
// 	//&out =
// 	cmd.Stdout = &out
// 	if err := cmd.Run(); err != nil {
// 		log.Printf("Cmd execute error!!info:\n", err.Error())
// 	}
// 	buf := make([]byte, 1024)
// 	//var strContent string = ""
// 	for {
// 		//read读取根本没有去取write进去的数据
// 		n, _ := out.Read(buf)
// 		if n == 0 {
// 			break
// 		}
// 		//strContent += string(buf[0:n])
// 		fmt.Println(n)
// 	}
// 	//fmt.Println(strContent)
// }
