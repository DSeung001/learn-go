package main

import (
	"fmt"
	"sync"
	"time"
)

func sqaure2(wg *sync.WaitGroup, ch chan int) {
	// range ch을 하는 순간 ch을 계속 기다리게 됨 => close(ch)를 통해 채널을 닫아줘야 함
	// range ch을 통해서 무한 대기하는 경우를 좀비 루틴, 고루틴릭이라고 부름
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go sqaure2(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널을 닫아줘야 range ch를 통해 채널을 계속 기다리는 것을 방지할 수 있음
	wg.Wait()
}
