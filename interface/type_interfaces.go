package main

//类型断言
import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
	//第二种switch
	switch tp := areaIntf.(type) {
	case *Square:
		fmt.Printf("%T", tp)
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

/*

switch areaIntf.(type) {
case *Square:
    // TODO
case *Circle:
    // TODO
...
default:
    // TODO
}














func classifier(items ...interface{}) {
    for i, x := range items {
        switch x.(type) {
        case bool:
            fmt.Printf("Param #%d is a bool\n", i)
        case float64:
            fmt.Printf("Param #%d is a float64\n", i)
        case int, int64:
            fmt.Printf("Param #%d is a int\n", i)
        case nil:
            fmt.Printf("Param #%d is a nil\n", i)
        case string:
            fmt.Printf("Param #%d is a string\n", i)
        default:
            fmt.Printf("Param #%d is unknown\n", i)
        }
    }
}


可以这样调用此方法：classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)













*/
