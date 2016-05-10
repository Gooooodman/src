package main

/*

异常
revover 捕获异常 但只能用在defer 中

*/
import (
	"fmt"
)

func Test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}

	}()
	divide(5, 0) //执行遇到错误 执行 defer
	fmt.Println("end if try")
}

func divide(a, b int) int {
	return a / b
}

func main() {
	Test()
}
