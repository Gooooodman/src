package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int    "编号"
	Name string "姓名"
	Sex  bool   "性别"
}

func main() {
	u := user{1, "张三", false}
	t := reflect.TypeOf(u)  // 类型
	v := reflect.ValueOf(u) //值
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s (%s = %v)\n", f.Tag, f.Name, v.Field(i).Interface())
		//fmt.Printf("%s (%s = %v)\n", f.Tag, f.Name, v.Field(i))  一样
	}

	name := reflect.ValueOf(u).FieldByName("Name")
	// 取值
	fmt.Println(name.Interface())

}
