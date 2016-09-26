-------------------------------------------------------------
//字符串转成[]byte
[]byte("Here is a string....")



-------------------------------------------------------------
type Integer int
func (a Integer) Less(b Integer) bool {
    return a < b
}
func (a *Integer) Add(b Integer) {
    *a += b
}

type LessAdder interface { 
    Less(b Integer) bool 
    Add(b Integer)
}

var a Integer = 1
var b1 LessAdder = &a //OK
var b2 LessAdder = a   //not OK

也就是说*Integer实现了接口LessAdder的所有方法，而Integer只实现了Less方法，所以不能赋值。


-------------------------------------------------------------


chan <- int 只写
<- chan int  只读



----------------------------------------------------------------

#Sscanf从字符串str扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数。返回成功扫描的条目个数和遇到的任何错误。
    var unit string
    var v float64
    str := "30.6C"
    fmt.Println(fmt.Sscanf(str, "%f%s", &v, &unit))
    fmt.Println(v)
    fmt.Println(unit)



----------------------------------------------
    var input string
    fmt.Scanln(&input)






fmt.Printf("%s\n",xx)
xx 可以是[]byte






//                包中的init会先调用

package salt

import (
    "fmt"
)

func init() {
    fmt.Println("salt init")
}

func Salttest() {
    fmt.Println("salttest")
}


//调用
package main

import (
    //"fmt"
    "salt"
)

func main() {
    // var S *salt.Server = new(salt.Server)
    // S.GetToken()
    // //fmt.Println(S)
    // ret := S.ListALLKey()
    // fmt.Println(ret)
    salt.Salttest()

}

//结果

salt init
salttest
