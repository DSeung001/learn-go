package ch1

import (
	"fmt"
	"sync"
	"time"
)

func Deadlock() {
	type value struct {
		mu       sync.Mutex
		resource int
	}

	var wg sync.WaitGroup
	process := func(v1, v2 *value) {
		defer func() {
			wg.Done()
			v1.mu.Unlock()
			v2.mu.Unlock()
		}()
		v1.mu.Lock()

		// 원할한 테스트를 위해 2초 대기
		time.Sleep(2 * time.Second)
		v2.mu.Lock()

		fmt.Println("sum=", v1.resource+v2.resource)
	}

	var a, b, c value
	wg.Add(3)

	// Process A 실행
	go process(&a, &c)
	// Process B 실행
	go process(&b, &a)
	// Process C 실행
	go process(&c, &b)

	wg.Wait()
}
