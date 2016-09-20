// test1.go project main.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	//var str string = ""
	buf := make([]byte, 1024)
	//f, _ := os.Open("doc.go")
	//for {
	//	n,_ := f.Read(buf)
	//	if n == 0{
	//		break
	//	}
	//	str += string(buf[0:n])
	//}
	file, _ := os.Create("doc1.go")
	w = file
	f, _ := os.Open("file_re.go")
	defer f.Close()
	defer file.Close()
	for {
		n, _ := f.Read(buf)
		fmt.Println(n)
		//fmt.Println(x)//x=1024
		if n == 0 {
			break
		}
		_, _ = w.Write(buf)
		//str += string(buf[0:n])
	}
	//fmt.Println(str)

	//x, _ := w.Write(buf)

}
