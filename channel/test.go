package main

import "fmt"

func foo() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

/*
chan <- int  写   可以close()
<- chan int  读   不能close()


*/

func sum(x, y int, c chan<- int) {

	c <- x + y
	close(c)
}

func main() {
	foo()
	c := make(chan int)
	go sum(24, 18, c)
	fmt.Println(<-c)
	// var input string
	// fmt.Scanln(&input)
}
