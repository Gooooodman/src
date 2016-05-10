package main

import (
	"fmt"
	//"time"
)

func Count(chs []chan int, ch chan int) {

	defer func() {

		ch <- 100
	}()

	for i := 0; i < 10; i++ {

		chs[i] = make(chan int)
		chs[i] <- 1
		close(chs[i])
	}

}

func main() {

	chs := make([]chan int, 10)
	ch := make(chan int)

	go Count(chs, ch)

	for val := range chs {
		fmt.Println(val)
	}

	<-ch

}
