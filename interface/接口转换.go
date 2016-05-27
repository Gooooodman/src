package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

type Student struct {
	People
	School string
}

type PeopleInfo interface {
	GetPeopleInfo()
}

type StudentInfo interface {
	GetPeopleInfo()
	GetStudentInfo()
}

func (p People) GetPeopleInfo() {
	fmt.Println(p)
}

//这里使用了*
func (s *Student) GetStudentInfo() {
	s.Name = "老王"
	fmt.Println(s)

}

func main() {
	//这里要使用&
	var is StudentInfo = &Student{People{"李明", 18}, "qinghua"}
	is.GetPeopleInfo()
	is.GetStudentInfo()

	//接口转换为了子集
	var ip PeopleInfo = is
	ip.GetPeopleInfo()

}
