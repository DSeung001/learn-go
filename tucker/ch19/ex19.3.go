package main

import "fmt"

type account2 struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account2) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account2) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account2) withdrawReturnValue(amount int) account2 {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account2 = &account2{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // 여전히 70

	var mainB account2 = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) // 50

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance) // 20
}
