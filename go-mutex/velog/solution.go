package velog

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 출처 : https://velog.io/@supssson/Go-Mutex%EB%8A%94-%EC%96%B4%EB%96%BB%EA%B2%8C-%EB%8F%99%EC%9E%91%ED%95%A0%EA%B9%8C

// Solution1 : 고루틴 동기성 문제 1번 임의의 값 사용
func Solution1() {
	var lock int = 1
	num := 0

	go func() {
		// Lock Acquire 과정
		for lock == 1 {
			lock -= 1
		}
		fmt.Println("Lock Acquired [1]")
		for i := 0; i < 1000000; i++ {
			num += 1
		}
		if lock == 0 {
			lock += 1
		}
		fmt.Println("Lock Released [1]")
	}()

	go func() {
		// Lock Acquire 과정
		for lock == 1 {
			lock -= 1
		}
		fmt.Println("Lock Acquired [2]")
		for i := 0; i < 1000000; i++ {
			num += 1
		}
		// Lock	Release 과정
		if lock == 0 {
			lock += 1
		}
		fmt.Println("Lock Released [2]")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("After goroutine", num, lock)
}

// Solution2 : 고루틴 동기성 문제 2번 Mutex 사용
func Solution2() {
	mu := sync.Mutex{}
	num := 0

	go func() {
		// Lock Acquire 과정
		mu.Lock()
		fmt.Println("Lock Acquired [1]")
		for i := 0; i < 1000000; i++ {
			num += 1
		}
		// Lock	Release 과정
		fmt.Println("Lock Released [1]")
		mu.Unlock()
	}()

	go func() {
		// Lock Acquire 과정
		mu.Lock()
		fmt.Println("Lock Acquired [2]")
		for i := 0; i < 1000000; i++ {
			num += 1
		}
		// Lock	Release 과정
		fmt.Println("Lock Released [2]")
		mu.Unlock()
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("After goroutine", num)
}

// Solution3 : 고루틴 동기성 문제 3번, mutex가 아닌 atomic 사용
func Solution3() {
	var mutex int32 = 1
	num := 0

	go func() {
		// Lock Acquire 과정
		for !atomic.CompareAndSwapInt32(&mutex, 1, 0) {
		}
		fmt.Println("Lock Acquired [1]")
		for i := 0; i < 1000000; i++ {
			num += 1
		}
		fmt.Println("Lock Released [1]")
		atomic.AddInt32(&mutex, 1)
	}()

	go func() {
		// Lock Acquire 과정
		for !atomic.CompareAndSwapInt32(&mutex, 1, 0) {
		}
		fmt.Println("Lock Acquired [2]")
		for i := 0; i < 1000000; i++ {
			num += 1
		}
		// Lock Release 과정
		fmt.Println("Lock Released [2]")
		atomic.AddInt32(&mutex, 1)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("After goroutine", num, mutex)
}
