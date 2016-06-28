package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		//fmt.Println("i: ", i)
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		//fmt.Println("in: ", i)
		//fmt.Println("prime: ", prime)
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
			//fmt.Println("out: ", <-out)
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		//fmt.Println("==========================================================================================: ", i)
		//fmt.Println("ch0: ", ch)
		prime := <-ch
		fmt.Println(i, " : ", prime)
		ch1 := make(chan int)
		//fmt.Println("ch1: ", ch1)
		go Filter(ch, ch1, prime)
		ch = ch1
		//fmt.Println("ch1 : ", <-ch)
		//fmt.Println("==========================================================================================: ", i)
	}

}
