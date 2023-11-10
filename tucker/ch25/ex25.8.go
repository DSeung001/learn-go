package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

// 컨텍스트는 작업 가능 시간, 작업 취소등의 조건을 지시할 수 있는 작업 명세서 역할을 합니다.
func main() {
	wg2.Add(1)
	// context.WithCancel은 2개를 반환 context와 cancel 함수
	// context.Background가 기본이고 이처럼 상위 컨텍스트를 감싸는 형태로 컨텍스트를 생성
	ctx, cancel := context.WithCancel(context.Background()) // 취소가 가능한 컨텍스트 생성
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() // 컨텍스트 취소

	wg2.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second) // 1초 간격 시그널
	for {
		select {
		case <-ctx.Done(): // 컨텍스트가 취소되면
			wg2.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
