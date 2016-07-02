package main

import (
//"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	var b1 LessAdder = &a //OK

	// 也就是说*Integer实现了接口LessAdder的所有方法，而Integer只实现了Less方法，所以不能赋值。
	var b2 LessAdder = a //not OK
}
