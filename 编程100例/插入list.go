/*
程序 39：插入元素

题目：有一个已经排好序的数组。现输入一个数，要求按原来的规律将它插入数组中。
1. 程序分析：首先判断此数是否大于最后一个数，然后再考虑插入中间的数的情况，插入后此元素之后的数，依次后移一个位置。
2.程序源代码：
*/

package main

import (
	"fmt"
)

func main() {
	var number int
	var array = [11]int{1, 4, 6, 9, 13, 16, 19, 28, 40, 100}
	fmt.Printf("insert a new number:")
	fmt.Scanf("%d", &number)
	array[10] = number
	for i := 10; i > 0; i-- {
		if array[i] < array[i-1] {
			array[i], array[i-1] = array[i-1], array[i]
		} else {
			break
		}
	}
	fmt.Println(array)
}
