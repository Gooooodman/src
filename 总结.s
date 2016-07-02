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














