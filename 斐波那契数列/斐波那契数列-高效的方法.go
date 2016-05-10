package main

import (
	"fmt"
	"time"
)

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 1), make(chan int, 1), make(chan int, 1)
	go func() {
		for {
			//fmt.Println("in:", in)
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 1)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

func main() {
	start := time.Now()
	x := fib()
	for i := 0; i < 41; i++ {
		fmt.Println(<-x)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
