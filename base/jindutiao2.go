package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	ch := make(chan int)
	ret := make(chan bool)
	go JDT(ch, ret)
	pinger(ch, ret)
	fmt.Printf("\rover...............................................................................")
}

func JDT(ch chan int, ret chan bool) {
	for {
		select {
		case c := <-ch:
			h := strings.Repeat("=", c) + strings.Repeat(" ", 49-c)
			fmt.Printf("\r[%s]%.0f%%", h, float64(c)/49*100)
		case <-ret:
			break
		}
	}

}

func pinger(c chan int, ret chan bool) {
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- i
	}
	ret <- true
}

/*    最初
func main() {
    for i := 0; i < 50; i++ {
        time.Sleep(100 * time.Millisecond)
        h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
        fmt.Printf("\r[%s]%.0f%%", h, float64(i)/49*100)
        os.Stdout.Sync()
    }
    fmt.Println("\nAll System Go!")
}
*/
