package main

import (
	"fmt"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func frist(out chan<- int) {
	go func() {
		out <- fib(16)
	}()
	go func() {
		out <- fib(35)
	}()
}

func main() {
	ch := make(chan int)
	go frist(ch)
	fmt.Println(<-ch)
	fmt.Println("=========================================")
	fmt.Println(fib(35))
	fmt.Println(fib(16))
}
