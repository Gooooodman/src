package main

import (
	"fmt"
	"time"
)

/*
 <-tick   接收

*/

func main() {
	// 创建一个 tick channel
	// 在 100 毫秒后会向 tick channel 中发送当前时间
	tick := time.Tick(100 * time.Millisecond)
	// 创建一个 boom channel
	// 在 500 毫秒后会向 boom channel 中发送当前时间
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
