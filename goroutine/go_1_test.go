package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func f(c chan int, n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		//amt := time.Duration(rand.Intn(250))
		//time.Sleep(time.Millisecond * amt)
	}
	//fmt.Println("完成")
	c <- 1
}
func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go f(c, i)

	}
	<-c
	var input string
	fmt.Scanln(&input)
	//time.Sleep(1 * time.Second)
}
