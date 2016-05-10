package main

import (
	"fmt"
)

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)

}
func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

type X struct{}

func (*X) test() {
	println("X.test")
}

func main() {
	d := Data{}
	p := &Data{55}
	dd := d
	var pp *Data
	pp = p

	fmt.Printf("Data: %p\n", p)
	d.ValueTest() // ValueTest(d)
	d.x = 1
	fmt.Println(d)
	dd.x = 11
	fmt.Println(d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
	p.x = 2
	fmt.Println(p)

	c := new(Data)
	c.x = 3
	fmt.Println(c)
	pp.x = 33
	fmt.Println(p)
	x := &X{}
	x.test()
	//(&p).test()   error  不能使用多级指针查询

}
