package main

import (
	"fmt"
	"link"
)

/*单向链表*/

func main() {
	var head *link.Node //定义head 类型
	stu1 := link.Node{link.Student{100, "李四"}, nil}
	stu2 := link.Node{link.Student{50, "王五"}, nil}
	stu3 := link.Node{link.Student{300, "赵六"}, nil}
	// stu4 := link.Node{link.Student{400, "张三"}, nil}
	// stu5 := link.Node{link.Student{500, "鲁智深"}, nil}
	// stu6 := link.Node{link.Student{30, "赵本山"}, nil}
	// stu7 := link.Node{link.Student{700, "李白"}, nil}
	head = head.Creat()
	//fmt.Println(head)
	head = stu1.Insert(head)

	head = stu2.Insert(head)
	fmt.Println(head)
	head = stu3.Insert(head)
	fmt.Println(head)
	// head = stu4.Insert(head)
	// fmt.Println(head)
	// head = stu5.Insert(head)
	// fmt.Println(head)
	// head = stu6.Insert(head)
	// fmt.Println(head)
	// head = stu7.Insert(head)
	// fmt.Println(head)
	head.PrintLink()
	// head = stu4.Delete(head)
	// head.PrintLink()
	// head = stu4.Insert(head)
	// head.PrintLink()
}
