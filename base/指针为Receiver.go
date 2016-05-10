package main

import (
	"fmt"
)

type coordinate struct {
	x int
	y int
}

func (recv *coordinate) swap() {
	recv.x, recv.y = recv.y, recv.x
	fmt.Println("swap: ", recv)
}

func main() {
	r1 := coordinate{3, 4}
	fmt.Println("开始: ", r1)
	//p = &r1
	r1.swap()
	fmt.Println("最后: ", r1)
}
