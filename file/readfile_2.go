package main

import (
	"fmt"
	"os"
)

/*

读整个文件

*/

func main() {
	file, err := os.Open("D:\\新建文本文档.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	//返回文件的信息
	/*
		type FileInfo interface {
		    Name() string       // 文件的名字（不含扩展名）
		    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		    Mode() FileMode     // 文件的模式位
		    ModTime() time.Time // 文件的修改时间
		    IsDir() bool        // 等价于Mode().IsDir()
		    Sys() interface{}   // 底层数据来源（可以返回nil）
		}
	*/
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	s, err := file.Read(bs)
	if err != nil {
		return
	}
	fmt.Println("字节数: ", s)
	fmt.Println("-------------------------------------------")
	str := string(bs)
	fmt.Println(str)
}
