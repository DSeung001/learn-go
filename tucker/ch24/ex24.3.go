package main

import (
	"fmt"
	"sync"
	"time"
)

// 동시성 문제 => 같은 자원을 여러 고루틴이 동시에 접근할 떄 생기는 문제
// 뮤텍스로 해결할 수 있음 => Lock 메서드를 통해 공유 자원에 접근을 막고 Unlock 메서드를 통해 접근을 허용

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock() // 뮤텍스를 확보하기 전 까지 아래 코드 실행 x
	defer mutex.Unlock()

	// 뮤텍스를 확보한 하나의 고루틴만이 아래 코드를 실행
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance Shoud not be negative value : %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
