package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	num := make(chan int, len(s))

	go func() {
		for i := range s {
			go test(i, num)
		}
	}()

loop:
	for {
		select {
		case size, ok := <-num:
			if !ok {
				break loop
			}
			fmt.Println(size)
		}
	}

}

var ch = make(chan struct{}, 5)

func test(i int, num chan<- int) {
	select {
	case ch <- struct{}{}:
	}

	defer func() { <-ch }()
	num <- i
	fmt.Println("---------->", i)
	time.Sleep(2 * time.Second)
}
