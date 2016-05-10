package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	var input string
	fmt.Scanln(&input)
}

/*

6
0
2
12
8
10
4
16
14
18


*/
