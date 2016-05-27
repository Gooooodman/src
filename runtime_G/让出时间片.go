package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			if i == 2 {
				runtime.Gosched()
			}
			fmt.Println("Goroutine1: ", i)
		}
	}()
	go func() {
		fmt.Println("Goroutine2")
	}()
	time.Sleep(5 * time.Second)
}

/*
Goroutine1:  0
Goroutine1:  1
Goroutine2
Goroutine1:  2
Goroutine1:  3
Goroutine1:  4
*/
