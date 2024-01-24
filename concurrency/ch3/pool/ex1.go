package pool

import (
	"fmt"
	"sync"
)

// Pool 자원이 없으 떄만 New로 생성, Get으로쓰고 Put으로 반환
func Ex1() {
	var numCalcsCreated int
	calcPool := sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// 4KB 풀 생성
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			/*
				이 메모리에서 뭔가 흥미롭지만 빠른 작업이 이루어진다고 가정하자
			*/
		}()

	}

	wg.Wait()
	// 30개로 생각보다 덜 생성됐다 => 30kb
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
