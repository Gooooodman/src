package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//使用全部cpu核心
	b := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(40) // 填充
	for i := 0; i < 40; i++ {
		go Count(&wg, i)
	}

	wg.Wait() //等待完成
	fmt.Println(time.Since(b))
}

func Count(wg *sync.WaitGroup, i int) {
	//ch <- 1
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(i, ": Counting", a)
	wg.Done() // 完成一次减一次
}
