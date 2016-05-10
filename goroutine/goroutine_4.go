package main

import (
	"fmt"
	"runtime"
	//"time"
)

func test(ch chan int) {
	defer func() {
		fmt.Println("in  defer!")
		ch <- 1
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 5 {
			runtime.Goexit()
		}

	}
}

/*

如果没有ch 阻塞 main 就直接退出了 英文使用了go
goroutine_3  没有时间等待就直接结束了

*/

func main() {
	ch := make(chan int)
	go test(ch)
	//time.Sleep(5 * time.Second) 这种不合理
	//var str string
	//fmt.Scan(&str)
	//fmt.Println(&str)
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("cpu核数: ", n)
	<-ch
}
