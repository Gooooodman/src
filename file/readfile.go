package main

import (
	"fmt"
	"os"
)

/*

读文件操作
每次读1024字节
*/
func ReadFile(strFileName string) (string, error) {
	f, err := os.Open(strFileName)
	if err != nil {
		fmt.Println("read file  err!")
		return "", err
	}

	defer f.Close()
	buf := make([]byte, 1024)
	var strContent string = ""
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		strContent += string(buf[0:n])
	}
	return strContent, nil
}

func main() {
	str, err := ReadFile("test.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)
}
