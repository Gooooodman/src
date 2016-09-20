package main

import (
	"fmt"
	//"time"
	//"runtime"
)

func Count(i int, ch chan int) {
	//ch <- 1
	//fmt.Println(i, ": Counting")
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(i, ": Counting", a)
	ch <- 1
	close(ch)
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	chs := make([]chan int, 80)
	for i := 0; i < 80; i++ {
		chs[i] = make(chan int)
		go Count(i, chs[i])
	}
	//ch1 := make(chan int)
	for _, ch := range chs {
		<-ch
	}
}

// 	// var input string
// 	// fmt.Scanln(&input)
// }

// func main() {
// 	chs := make([]chan int, 10)
//     ch1 := make(chan int)
// 	for i := 0; i < 10; i++ {
// 		chs[i] = make(chan int)
// 		go Count(i, chs[i])
// 	}

// 	for _, ch := range chs {
// 		<-ch
// 	}
//     <-ch1
// 	var input string
// 	fmt.Scanln(&input)
// }

// func test(ch chan int) {
//     defer func() {
//         fmt.Println("in  defer!")
//         ch <- 1
//     }()
//     for i := 0; i < 10; i++ {
//         fmt.Println(i)
//         if i > 5 {
//             runtime.Goexit()
//         }

//     }
// }

//第二种

// func Count(ch chan int, chs []chan int) {
// 	defer func() {
// 		fmt.Println("完成....")
// 		close(ch)
// 	}()
// 	fmt.Println("counting...")
// 	for i := 0; i < 10; i++ {
// 		chs[i] = make(chan int)
// 		close(chs[i])
// 		fmt.Println(i, ": Counting")
// 	}
// }

// func main() {
// 	chs := make([]chan int, 10)
// 	ch := make(chan int)
// 	go Count(ch, chs)
// 	// for _, c := range chs {
// 	//  <-c
// 	// }

// 	<-ch

// 	time.Sleep(time.Second * 1)
// }

//第三中

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	//使用全部cpu核心
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	wg := sync.WaitGroup{}
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go GO(&wg, i)
// 	}

// 	wg.Wait()
// }

// func GO(wg *sync.WaitGroup, index int) {
// 	a := 1
// 	for i := 0; i < 1000000000; i++ {
// 		a += i
// 	}
// 	fmt.Println(index, a)
// 	//完成一次
// 	wg.Done()
// }
