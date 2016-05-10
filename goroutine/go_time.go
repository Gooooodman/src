package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//当前时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	/*
	  2016-01-09 14:49:00.846658 +0800 CST
	  2016-01-09 14:49:00
	*/

	d, err := time.Parse("01-02-2006", "01-09-2016")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)

}
