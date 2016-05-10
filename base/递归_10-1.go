package main

import (
	"fmt"
	"time"
)

func main() {
	//result := 0
	start := time.Now()
	// for i := 10; i >= 1; i-- {
	// 	result = fibonacci(i)
	// 	fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	// }
	fibonacci_10(10)
	fmt.Println("-----------------------------------------------------")
	fibonacci_1(1)
	end := time.Now()
	delta := end.Sub(start) // 计算函数时间
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

//10 ------------------- 1
func fibonacci_10(n int) {
	if n >= 0 {
		fmt.Println(n)
		//ret = n
		fibonacci_10(n - 1)
	}
	return
}

// 1 ------------------ 10
func fibonacci_1(n int) {
	if n <= 10 {
		fmt.Println(n)
		//ret = n
		fibonacci_1(n + 1)
	}
	return
}
