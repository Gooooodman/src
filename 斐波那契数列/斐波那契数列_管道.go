package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			//fmt.Println("y: ", y)
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			//c 接收一次 算一次i
			fmt.Println(i, <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
