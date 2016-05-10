package main

import (
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("test.go")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		//n 是读取的字节数
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
