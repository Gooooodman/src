package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// f, err := os.OpenFile("c:\\", os.O_RDONLY, 0666)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// arrfile, err1 := f.Readdir(0)
	// if err1 != nil {
	// 	fmt.Println(err1.Error())
	// 	return
	// }
	//
	//封装了f.Readdir()
	arrfile, err := ioutil.ReadDir("c:\\")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	check(arrfile)
	//读目录下的文件
	for k, v := range arrfile {
		fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
	}
}

//判断dir 类型
func check(v interface{}) {
	if _, ok := v.([]os.FileInfo); ok {
		fmt.Println("ok")
	}

}
