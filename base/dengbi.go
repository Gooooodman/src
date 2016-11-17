package main

import (
	"fmt"
	"math"
)

const (
	a = 20 // 第一位
	n = 8  // 连续10
	q = 2  // 平方

)

func main() {
	for i := 1; i <= 8; i++ {
		fmt.Printf("%v -----> %v\n", i, a*math.Pow(2, float64(i-1)))
	}
	fmt.Println("-----------------------------------------")
	fmt.Printf("%v 次总投入: %v\n", n, a*(1-math.Pow(q, n))/(1-q))
	fmt.Printf("%d 次需投入: %v\n", n+1, a*math.Pow(2, float64(n)))
}
