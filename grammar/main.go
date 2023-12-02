package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 지금 나는 CPU는 코어 8개고 논리 프로세스는 16개로 즉 16개의 스레드를 동시에 실행 가능
	fmt.Println(runtime.NumCPU())
}
