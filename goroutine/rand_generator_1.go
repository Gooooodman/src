package main

import (
	"fmt"
	"math/rand"
)

// 函数 rand_generator_2，返回 通道(Channel)
func rand_generator_2() chan int {
	// 创建通道
	out := make(chan int)
	// 创建协程
	go func() {
		for {
			//向通道内写入数据，如果无人读取会等待
			out <- rand.Int()
		}
	}()
	return out
}

// 函数 rand_generator_1 ，返回 int
func rand_generator_1() int {
	return rand.Int()
}

func rand_generator_3() chan int {
	// 创建两个随机数生成器服务
	rand_generator_1 := rand_generator_2()
	rand_generator_2 := rand_generator_2()

	//创建通道
	out := make(chan int)

	//创建协程
	go func() {
		for {
			//读取生成器1中的数据，整合
			out <- <-rand_generator_1
		}
	}()
	go func() {
		for {
			//读取生成器2中的数据，整合
			out <- <-rand_generator_2
		}
	}()
	return out
}

func main() {
	// 生成随机数作为一个服务
	rand_service_handler := rand_generator_3()
	// 从服务中读取随机数并打印
	fmt.Printf("%d\n", <-rand_service_handler)

	fmt.Println(rand_generator_1())
}
