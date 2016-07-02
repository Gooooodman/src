package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		//fmt.Println("i: ", i)
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func pch(ch <-chan int, out chan<- int) {
	for {
		i := <-ch
		fmt.Println("i: ", i)
		out <- i
	}
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		a := <-ch
		fmt.Println("a: ", a)
		go pch(ch, ch1)
		ch = ch1
	}

	var input string
	fmt.Scanln(&input)
}
