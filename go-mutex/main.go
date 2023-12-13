package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	count := 20
	account := &Account{Balance: 0}

	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			for {
				AccountRecords(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func AccountRecords(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value :%d", account.Balance))
	}

	account.Balance += 100
	time.Sleep(time.Millisecond)
	account.Balance -= 100
}
