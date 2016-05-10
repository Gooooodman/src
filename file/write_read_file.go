package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("D:\\新建文本文档.txt", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	f.WriteString("go go  go")
	buf := make([]byte, 1024)
	var str string
	f.Seek(0, os.SEEK_SET) //重置指针
	for {
		n, ferr := f.Read(buf)
		if ferr != nil && ferr != io.EOF {
			fmt.Println(ferr.Error())
			break
		}
		if n == 0 {
			break
		}
		fmt.Println(n)
		str += string(buf[0:n])
	}
	fmt.Println(str)
	f.Close()

}
