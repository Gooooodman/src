/*

生产者与消费者

*/

package main

import "fmt"

// func producer(c chan int) {
// 	defer close(c)
// 	for i := 0; i < 10; i++ {
// 		c <- i //接收一个阻塞
// 	}
// }

// func consumer(c, f chan int) {
// 	// for {
// 	// 	if v, ok := <-c; ok {
// 	// 		fmt.Println(v)
// 	// 	} else {
// 	// 		break
// 	// 	}
// 	// }
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// 	f <- 1
// }

/*
可以将channel 指定为单向通信。
<-chan int  仅能接收
chan <- int 仅能发送
*/

func producer(c chan<- int) {
	//发送的c 进行了close(c)
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i //发送一个阻塞
	}
}

func consumer(c <-chan int, f chan<- int) {
	//           接收           发送
	// for {
	//  if v, ok := <-c; ok {
	//      fmt.Println(v)
	//  } else {
	//      break
	//  }
	// }
	for v := range c {
		fmt.Println(v)
	}
	f <- 1
}

func main() {
	buf := make(chan int)
	flg := make(chan int)
	go producer(buf)
	go consumer(buf, flg)
	<-flg
}
