package main

import (
	"fmt"
)

func main() {
	m := map[int]struct { //map 键位int, 值为struct
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}

	fmt.Println(m[1].name)

	n := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	if v, ok := n["a"]; ok {
		fmt.Println(v)
		fmt.Println(ok) //true
	}

	for k, v := range n {
		fmt.Println(k, v)
	}

}
