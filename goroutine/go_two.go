package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	sige := make(chan byte, 2)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("channel close")
		sige <- 0
	}()

	go func() {
		for {
			e, ok := <-ch
			fmt.Printf("%d (%v)\n", e, ok)
			if !ok {
				break
			}
			//time.Sleep(2 * time.Second)
		}
		fmt.Println("Done")
		sige <- 1
	}()

	<-sige
	<-sige
}
