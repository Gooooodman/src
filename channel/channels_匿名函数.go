package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //  匿名函数要关闭
	}()

	// for value := range ch {
	// 	fmt.Println(value)
	// }

	for {
		if value, ok := <-ch; ok {
			fmt.Println(value)
		} else {
			break
		}
	}
}
