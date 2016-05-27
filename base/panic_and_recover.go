package main

import (
	"fmt"
)

func main() {
	//必须要先声明defer  否则捕获不到panic
	defer func() {
		fmt.Println("函数 defer 开始运行...")
		if err := recover(); err != nil {
			//这里的err就是panic转入的内容
			fmt.Println("程序异常退出: ", err)
		} else {
			fmt.Println("程序正常退出")
		}

	}()
	f(101)
}

func f(a int) {
	fmt.Println("函数 f 开始执行...")
	if a > 100 {
		panic("参数值超出范围")

	} else {
		fmt.Println("函数f结束")
	}
}
