package main

import (
	"fmt"
	"time"
)

func process() {
	// 작업 완료 여부를 전달할 채널
	done := make(chan error)
	// 작업 수행
	go doSomething(done)
	// 타입 아웃 채널
	timeOut := time.After(3 * time.Second)

	select {
	case err := <-done: // 작업 완료
		if err != nil {
			fmt.Println("Status Request Timeout")
		} else {
			fmt.Println("Success")
		}
	case <-timeOut: // 3초 후 타임아웃
		fmt.Println("Status Request Timeout")
	}
}

func doSomething(done chan<- error) {
	fmt.Println("Processing request...")
	select {
	case <-time.After(5 * time.Second): // 5초 걸리는 작업
		fmt.Println("Task completed successfully")
		done <- nil // 작업 성공 신호 전송
	}
}

func main() {
	process()
}
