package main

import (
	"fmt"
)

//这个函数很有意思  接收一个参数 返回bool
type MyFunType func(int) bool

func IsBigThan5(n int) bool {
	return n > 5
}

func Display(arr []int, f MyFunType) {
	for _, v := range arr {
		if f(v) {
			fmt.Println(v)
		}
	}
}

func main() {
	arr := []int{1, 2, 4, 5, 6, 7, 3, 9} //切片
	Display(arr, IsBigThan5)
	// arr[0] = 10
	// fmt.Println(arr)
}
