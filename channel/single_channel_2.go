package main

import (
	"fmt"
)

func Recv(ch <-chan int, lock chan<- bool) {
	for value := range ch {
		fmt.Println(value)
	}
	lock <- true
	close(lock)
}

func Send(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2) //双向通道转换单向通道                   异步通道 达到消息队列的效果  传输大量数据
	lock := make(chan bool)
	go Recv(ch, lock) // 只能从ch接收
	go Send(ch)       //只能从ch发送
	<-lock
}

/*

ch := make(chan int)  不加buffer 同步

ch := make(chan int, 2)  加了buffer 异步




*/
