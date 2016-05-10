package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func main() {
	u := User{1, "tom"}
	var i interface{} = &u //数据制作持有的目标对象的只读复制品,复制完整对象或指针
	u.id = 2
	u.name = "jack"
	fmt.Printf("%v\n", u)
	i.(*User).id = 8
	fmt.Printf("%v\n", i.(*User))
	fmt.Printf("%v\n", i)

}

// func main() {
// 	//接口转型返回临时对象，只有使用指针才能修改其状态
// 	u := User{1, "Tom"}
// 	var vi, pi interface{} = u, &u
// 	// vi.(User).name = "Jack"         // Error: cannot assign to vi.(User).name
// 	pi.(*User).name = "Jack"
// 	fmt.Printf("%v\n", vi.(User))
// 	fmt.Printf("%v\n", pi.(*User))
// }
