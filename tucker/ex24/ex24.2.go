package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup을 통해 서브 고루틴을 기다릴 수 있음
var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done() // 작업 완료
}

func main() {
	const cnt int = 10
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go SumAtoB(1, 1000000000)
	}
	wg.Wait()
}
