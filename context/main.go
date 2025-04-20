package main

import (
	"context"
	"fmt"
	"time"
)

func process() {
	// 3초 후 종료하는 컨텍스트
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 작업 수행
	err := doSomething(ctx)

	// 타임 아웃에 따른 에러 반환
	if err != nil {
		fmt.Println("Status Request Timeout")
	} else {
		// 정상 처리
		fmt.Println("Success")
	}
}

func doSomething(ctx context.Context) error {
	fmt.Println("Processing request...")

	select {
	// 5초 작업
	// select 문과 같이 사용해야지, 컨텍스트 취소를 감지 가능
	case <-time.After(5 * time.Second):
		fmt.Printf("Task completed successfully")
		return nil
	case <-ctx.Done(): // context 에 종료 신호가 올 경우
		fmt.Println("Task aborted due to timeout:", ctx.Err())
		return fmt.Errorf("Request Timeout")
	}
}

func main() {
	process()
}
