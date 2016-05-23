package main

import (
	"fmt"
)

func main() {
	var array = [11]int{1, 4, 6, 9, 13, 16, 19, 28, 40, 100, 124}
	count := len(array)
	fmt.Println(count / 2)
	for i := 0; i < count/2; i++ {
		array[i], array[count-i-1] = array[count-i-1], array[i]
	}
	fmt.Println(array)
}
