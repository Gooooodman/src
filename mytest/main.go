package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Id     int
	Name   string
	Scores string
}

type StudentManager struct {
	student []Student
}

func NewStudentManager() *StudentManager {
	return &StudentManager{make([]Student, 0)}
}

func (s *StudentManager) Len() int {
	return len(s.student)
}

func (s *StudentManager) Add(stu *Student) {
	s.student = append(s.student, *stu)
}

func (s *StudentManager) Get(index int) (stu *Student, err error) {
	if index < 0 || index >= len(s.student) {
		return nil, errors.New("超出列表")
	}
	return &s.student[index], nil
}

//var id int = 0

func main() {
	var stu *StudentManager
	stu = NewStudentManager()
	fmt.Println(`
        Enter add  name scores //eg : add  lpx 30  增加一项lpx 30分
        Enter list       //列出
        Enter q|e        //退出
        `)
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		//fmt.Println(line)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "add" {
			handleCommands(stu, tokens)
		} else if tokens[0] == "list" {
			handleCommands(stu, tokens)
		} else {
			fmt.Println("Unrecognized command")

		}
	}

}

func handleCommands(stu *StudentManager, tokens []string) {
	// if len(tokens) < 2 {
	// 	fmt.Println("参数不足")
	// 	return
	// }
	id := 0
	switch tokens[0] {
	case "add":

		if len(tokens) == 3 {
			id++
			stu.Add(&Student{id, tokens[1], tokens[2]})
			fmt.Println("OK")
		} else {
			fmt.Println("用法: add name scores")
		}
	case "list":
		fmt.Println("序号   名字   分数")
		for i := 0; i < stu.Len(); i++ {
			//s, _ := stu.Get(i)
			s := stu.student[i]
			fmt.Printf("%-4d %-10s %-3s\n", s.Id, s.Name, s.Scores)
		}
	default:

		fmt.Println("Unrecognized lib command:", tokens[0])

	}

}
