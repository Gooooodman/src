package main

import (
	"fmt"
	"runtime"
	//"time"
)

//runtime.Gosched() 释放cpu权限

func SayHello() {
	// defer func() {
	// 	fmt.Println("ok")
	// 	c1 <- 1
	// }()
	for i := 0; i < 10; i++ {
		fmt.Print(i, " Hello ")
		runtime.Gosched()
	}
}

func SayWorld(c chan int) {
	defer func() {
		fmt.Println("ok")
		c <- 1
	}()
	for i := 0; i < 10; i++ {
		fmt.Println("World!", i)
		runtime.Gosched()
	}
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan int)
	//c1 := make(chan int)
	go SayHello()
	go SayWorld(c)
	//<-c1
	<-c

	//time.Sleep(1 * time.Second)
}
