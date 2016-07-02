package main

import (
	"fmt"
)

func main() {
	ch := make([]chan bool, 20)
	for i := 0; i < 20; i++ {
		ch[i] = make(chan bool)
		// chan <- bool  只写
		go func(i int, ch chan<- bool) {
			time.Sleep(5 * time.Second)
			fmt.Println(i)
			ch <- true
			// 写 关闭
			close(ch)
		}(i, ch[i])
	}
	//time.Sleep(10 * time.Second)
	for _, c := range ch {
		<-c
	}
}
