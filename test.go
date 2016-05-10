// package main

// import (
// 	"fmt"
// 	"net"
// 	"network"
// )

// // func main() {
// // 	s := make([]int, 0, 1)
// // 	c := cap(s)

// // 	for i := 0; i < 50; i++ {
// // 		s = append(s, i) //append 2倍涨容量
// // 		// fmt.Printf("i: %d ,s:%v", i, s)
// // 		if n := cap(s); n > c {
// // 			fmt.Printf("i:%d,s:%v,cap(s):%d\n", i, s, n)
// // 			fmt.Printf("cap: %d - > %d\n", c, n)
// // 			c = n
// // 		}
// // 	}
// // }

// //带缓冲
// // func main() {
// // 	c := make(chan int, 2)
// // 	c <- 1
// // 	c <- 2
// // 	fmt.Println(<-c)
// // 	fmt.Println(<-c)
// // }

// func main() {
// 	//ip := network.GetIp()
// 	ip := network.GetLookup("www.baidu.com")
// 	fmt.Println(ip)

// 	ips, _ := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
// 	fmt.Println(ips)

// }

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type R struct {
	name string
	age  int
}

type i interface {
	Start()
}

func (s *R) Start() {
	fmt.Println("start")
	fmt.Println(s.age)
}

func t(i interface{}) (ty string) {
	switch i.(type) {
	case string:
		fmt.Println("string")
		ty = "string"
	case R:
		fmt.Println("type")
	default:
		fmt.Println("no")
	}
	return
}

func main() {
	//fmt.Println(split(17))
	/*
		    for {
				fmt.Println("for")
			}
	*/
	//mystring := "hello world"
	//fmt.Println([]byte(mystring))
	//fmt.Println(string([]byte(mystring)))
	// R.name = "lpx"
	// R.age = 10

	//fmt.Println(R)
	r1 := new(R)
	r1.age = 21
	r1.name = "lpl"
	r1.Start()
	var p i = r1
	//s := "hello"
	// var p1 Person
	// p2 := new(Person)
	// fmt.Printf("%v\n", p1)
	// fmt.Printf("%v\n", p2)
	//r := &R{"LPX", 22}
	p.Start()
	//ty := t(s)
	//fmt.Println(ty)
	r1.age = 22
	fmt.Println("type:", reflect.TypeOf(r1))
	//r.age = 25
	//fmt.Println(r)
	p.Start()

	var p1 i = &R{"lupuxiao", 23}
	p1.Start()
	p.Start()
	f, _ := os.Open("go_1.go")
	file, _ := os.Create("go_1_test.go")
	if _, err := io.Copy(file, f); err != nil {
		fmt.Println("err", f, file)
		log.Fatal(err)
	}

	defer f.Close()
	defer file.Close()

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[1:4:5]
	fmt.Println(slice)

	s := strings.Split("abc abc", " ")
	fmt.Println(s, len(s))

	dir, _ := ioutil.ReadDir(".")
	// if _, ok := dir.(os.FileInfo); ok {
	// 	fmt.Println("dir ok")
	// }
	check(dir)
	//var dir interface{}  err

	fmt.Println("type:", reflect.TypeOf(dir))
	//fmt.Println(dir.Name())
	// 循环所有目录下的文件
	for _, v := range dir {
		fmt.Println(v.Name())
	}

	ch1 := make(chan int, 1)
	//ch1 <- 1
	//ch1 <- 2
	out := Mych1(ch1)
	go func() {
		for {
			ch1 <- 1
			fmt.Println(out)
		}
	}()

	var input string
	fmt.Scanln(&input)

}

func check(v interface{}) {
	if _, ok := v.([]os.FileInfo); ok {
		fmt.Println("ok")
	}

}

func Mych1(ch1 <-chan int) int {
	go func() {
		for {
			fmt.Println(<-ch1)
		}
	}()
	return 0
}
