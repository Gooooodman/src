package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	//var c chan string = make(chan string)
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	//time.Sleep(20 * time.Second)
	//cmd  输入就退出
	var input string
	fmt.Scanln(&input)
}
