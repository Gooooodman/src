package main

import (
	"fmt"
	"github.com/robfig/cron"
	//"time"
)

/*
结构体成员是一个函数

*/

type Test struct {
	f func(string) string
}

// func run() *Test {
// 	test := Newtest()
// 	return test
// }

//生成f
func Newtest() *Test {
	test := &Test{}
	test.f = func(name string) string {
		title := fmt.Sprintf("you name is : %s", name)
		return title
	}
	return test
}

func dis() {
	t := Newtest()
	fmt.Println(t.f("lupuxiao"))

}

//计划任务
func main() {
	spec := "*, *, *, *, *, *"
	c := cron.New()
	c.Start()
	c.AddFunc(spec, dis)
	select {}
}
