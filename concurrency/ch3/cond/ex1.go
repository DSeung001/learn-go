package cond

import (
	"fmt"
	"sync"
	"time"
)

func Ex1() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue")
		c.L.Unlock()
		// cond 타입의 Signal 메서드를 호출하면 대기중인 고루틴 중 하나가 풀린다.
		// => 시그널을 보내줘~ 느낌
		c.Signal()
		// broadcast는 대기중인 모든 고루틴을 풀어준다.
	}

	// cond로 풀릴떄까지 기다리는 코드를 짤 수 있다.
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct {
		}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
