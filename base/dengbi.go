package main

import (
	"fmt"
	"math"
)

const (
	a = 2  // 第一位
	q = 2  // 平方
	n = 15 // 连续10
)

func main() {

	fmt.Println(a * (1 - math.Pow(q, n)) / (1 - q))
	fmt.Println(math.Pow(2, 16))
}
