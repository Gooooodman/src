package main

import (
	"fmt"
	"time"
	//"strconv"
)

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	const n = 45
// 	fibN := fib(n) // slow
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}



func main() {
	//go spinner(100 * time.Millisecond)
	//const n = 45
	out:=make(chan string)
	l:=[]int{15,30,5}
	for _,n := range l{
		go func(n int){
			fibN := fib(n) // slow
			
			//out <-  fmt.Sprintf(strconv.Itoa(n-1)+"---->"+strconv.Itoa(fibN))
			// fmt.Sprintf 具有格式化
			out <-  fmt.Sprintf("\rFibonacci(%d) = %d", n-1, fibN)

		}(n+1)

	}

	for  range l{
		//if m := <-out{
			fmt.Println(<-out)
		//}
		
	}		
	//close(out)
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)



}



