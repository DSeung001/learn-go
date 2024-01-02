package ch1

import (
	"fmt"
	"sync"
)

func RaceCondition() {
	var num int
	go func() {
		num++
	}()
	if num == 0 {
		fmt.Printf("The value is %d.\n", num)
	}
}

func RaceConditionSolution() {
	var memoryAccess sync.Mutex
	var num int

	go func() {
		// memoryAccess.Lock()을 호출하면 Unlock()을 호출하기 전까지 다른 곳에서 Lock()을 호출할 수 없음
		memoryAccess.Lock()
		num++
		memoryAccess.Unlock()
	}()

	// Unlock()을 호출할 때 까지 기다림
	memoryAccess.Lock()
	if num == 0 {
		fmt.Printf("The value is %d.\n", num)
	} else {
		fmt.Printf("The value is %d.\n", num)
	}
	memoryAccess.Unlock()
}
