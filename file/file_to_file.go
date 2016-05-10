// test1.go project main.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	var str string = ""
	buf := make([]byte, 1024)
	//f, _ := os.Open("doc.go")
	//for {
	//	n,_ := f.Read(buf)
	//	if n == 0{
	//		break
	//	}
	//	str += string(buf[0:n])
	//}
	file, _ := os.Create("doc_1.go")
	w = file
	f, _ := os.Open("doc.go")
	for {
		n, _ := f.Read(buf)
		fmt.Println(n)
		_, _ = w.Write(buf)
		//fmt.Println(x)//x=1024
		if n == 0 {
			break
		}
		str += string(buf[0:n])
	}
	fmt.Println(str)

	//x, _ := w.Write(buf)

}
