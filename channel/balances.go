package main

import (
	"fmt"
)

var deposites = make(chan int)
var balances = make(chan int)

func Deposite(amount int) {
	deposites <- amount
}

func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposites:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go teller()
	Deposite(100)
	fmt.Println(Balance())
}
