package main

import (
	"fmt"
	"math/rand"
)

func Test(ch chan int) {
	fmt.Println("Go...")
	ch <- rand.Int()
	close(ch) //发送的时候关闭

}

func main() {
	chs := make([]chan int, 30)
	for i := 0; i < 30; i++ {
		chs[i] = make(chan int, 10) //满10 才阻塞         异步通道 达到消息队列的效果  传输大量数据
		go Test(chs[i])

	}
	for _, ch := range chs {
		value := <-ch
		fmt.Println(value)
	}
}
