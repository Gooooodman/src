package main

import (
	"fmt"
)

/*
接口类型变量李明可以存储任意类型的数值
如何反向知道实际保存的是哪一种类型的对象呢?

两种常用的方法进行判断:
Comma-ok 断言            格式：  value,ok = element.(T)
Switch 测试

*/

type People struct {
	Name string
	Age  int
}

type Tester interface{} //空接口库存在任意类型

func main() {
	people := People{"张三", 20}
	it := make([]Tester, 4) //使用空接口很巧妙
	it[0], it[1], it[2], it[3] = 1, "Hello", people, true

	for index, element := range it {
		fmt.Println("################# Comma-ok 断言 方式 #################")
		if value, ok := element.(int); ok {
			fmt.Printf("it[%d] is an int. value = %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("it[%d] is a string. value = %s\n", index, value)
		} else if value, ok := element.(People); ok {
			fmt.Printf("it[%d] is a People. value = %v\n", index, value)
		} else if value, ok := element.(bool); ok {
			fmt.Printf("it[%d] is a bool. value = %v\n", index, value)
		} else {
			fmt.Printf("it[%d] is an unknown type\n", index)
		}

		fmt.Println("################# switch 方式 #################")
		switch value := element.(type) {
		case int:
			fmt.Printf("it[%d] is an int. value = %d\n", index, value)
		case string:
			fmt.Printf("it[%d] is a string. value = %s\n", index, value)
		case People:
			fmt.Printf("it[%d] is a People. value = %v\n", index, value)
		case bool:
			fmt.Printf("it[%d] is a bool. value = %v\n", index, value)
		default:
			fmt.Printf("it[%d] is an unknown type\n", index)
		}

	}
}
