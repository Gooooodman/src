package main

import (
	"fmt"
)

func main() {
	f := closures(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

//返回一个函数极其该函数返回的值
func closures(x int) func(int) int {
	return func(y int) int {
		// 内部函数的引用的参数超出了它的作用范围 x
		return x + y
	}
}
