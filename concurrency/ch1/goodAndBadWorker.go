package ch1

import (
	"fmt"
	"sync"
	"time"
)

func GoodAndBadWorker() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			// greedWorker의 는 한번 돌 때 마다 sharedLock을 1번 잠그고 풀어줌
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("욕심쟁이 루프는 %v번 돌았음\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			// politeWorker의 루프는 한번 돌 때 마다 sharedLock을 3번 잠그고 풀어줌
			// lock과 unlock 함수의 시간차이는 없음
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("공손한 루프는 %v번 돌았음\n", count)
	}

	wg.Add(2)

	// greed가 많은 이유는
	// - polite가 한번 돌때 거의 2~3배 더 돌기 때문
	// - polite는 필요할 때만 사용한 반면 greed는 필요 없을 때도 사용하기 때문
	// - greed는 polite의 unlock이 끝나자말자 바로 뺏기 때문
	go greedWorker()
	go politeWorker()
	wg.Wait()
}
