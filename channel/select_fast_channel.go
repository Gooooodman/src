package main

import (
    "fmt"
    // "time"
)




func fibonacci(n int) (res int) {
    if n <= 1 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    return
}


/*

那么channel的缓存队列将不是满的也不是空的（图8.4），因此对该channel执行的发送或接收操作都不会发送阻塞。通过这种方式，channel的缓存队列解耦了接收和发送的goroutine


一种很不错的机制,提取最快的

*/
func main() {
    responses := make(chan int,3)
    go func() {responses <- fibonacci(30)}()
    go func() {responses <- fibonacci(40)}()
    go func() {responses <- fibonacci(20)}() 
    fmt.Println(<-responses) 
}