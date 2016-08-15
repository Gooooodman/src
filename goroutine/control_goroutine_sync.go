package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	num := make(chan int)

	//-------------------   同步 ------------------------ //
	// go func() {
	// 	for i := range s {
	// 		test(i, num)
	// 	}
	// 	close(num)
	// }()

	var n sync.WaitGroup
	//go func() {
	for i := range s {
		n.Add(1)
		go test(i, &n, num)
	}
	//close(num)
	//}()

	go func() {
		n.Wait()
		close(num)
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

func test(i int, n *sync.WaitGroup, num chan<- int) {
	defer n.Done()
	select {
	case ch <- struct{}{}:
	}

	defer func() { <-ch }()
	num <- i
	fmt.Println("---------->", i)
	time.Sleep(2 * time.Second)
}
