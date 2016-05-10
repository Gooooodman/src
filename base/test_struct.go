package main

import (
	"fmt"
)

/*

值类型

面相对象三大特性里：go 仅支持封装
没有class关键字, ，没有继承与多态


*/
type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

type File struct {
	name string
	size int
	attr struct {
		perm  int
		owner int
	}
}

type Resource struct {
	id   int
	name string
}

type Classift struct {
	id int
}

type User struct {
	Resource // 与Classift id 有重复
	Classift
	name string
}

func main() {
	n1 := Node{
		id:   1,
		data: nil,
	}

	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}

	fmt.Println(n1)
	fmt.Println(n2)

	//支持匿名结构

	f := File{
		name: "test.txt",
		size: 1025,
	}

	//第一种赋值方式
	f.attr.owner = 0755
	f.attr.perm = 2
	//第二种
	fmt.Println(f)
	var attr = struct {
		perm  int
		owner int
	}{2, 0755} //要写在一行

	f.attr = attr

	fmt.Println(f)

	u := User{
		Resource{1, "People"},
		Classift{100},
		"JACK",
	}

	fmt.Println(u.name)
	fmt.Println(u.Resource.name)
	fmt.Println(u.Classift.id)

	var user User
	user.Classift.id = 100
	user.Resource.id = 200
	user.name = "joy"
	user.Resource.name = "R"
	fmt.Println(user)

}
