package main

import "fmt"

type account struct {
	balance int
}

func withdraw(a *account, val int) { // 일반함수
	a.balance -= val
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}

	withdraw(a, 30)

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)
}
