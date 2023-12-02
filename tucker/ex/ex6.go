package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go printTick(ctx)
	time.Sleep(5 * time.Second)

	cancel() // cancel을 실행함으로써 ctx에 done 신호가 들어감

	wg.Wait()
}

func printTick(ctx context.Context) {
	// 초당 한번씩 이벤트 생성
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-ctx.Done(): // ctx에 done 신호가 들어오면
			fmt.Println("cancel() 실행으로 ctx.Done()이 실행됨")
			wg.Done() // waitGroup을 종료
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
