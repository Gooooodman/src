package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//使用全部cpu核心
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(40)
	for i := 0; i < 40; i++ {
		go Count(&wg, i)
	}

	wg.Wait()
}

func Count(wg *sync.WaitGroup, i int) {
	//ch <- 1
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(i, ": Counting", a)
	wg.Done()
}
