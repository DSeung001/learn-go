package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	iterations     = 500000 // 반복 횟수
	goroutineCount = 4      // 고루틴 수
)

// sync.Pool 기본 문법
// New 필드에 객체가 없을 때 생성할 함수를 지정
var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// workWithoutPool은 매 반복마다 샘 bytes.Buffer를 생성
func workWithoutPool(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		b := new(bytes.Buffer)
		b.WriteString("some text")
		_ = b.String()
	}
}

// workWithPool은 sync.Pool을 사용해 bytes.Buffer를 재사용
func workWithPool(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		// 풀에서 버퍼를 가져옴
		b := bufPool.Get().(*bytes.Buffer)
		b.Reset()
		b.WriteString("some text")
		_ = b.String() // 사용 예시
		// 사용 후 풀에 다시 반납
		bufPool.Put(b)
	}
}

// runTest로 Pool 쓰는 함수와 쓰지 않는 방식의 소요 시간, 메모리 사용량, GC 실행 횟수 비교
func runTest(withPool bool) {
	var wg sync.WaitGroup

	// 측정을 위한 초기 메모리 상태 읽기
	var memBefore, memAfter runtime.MemStats
	runtime.GC() // GC 실행해 초기 상태 정리
	runtime.ReadMemStats(&memBefore)

	startTime := time.Now()

	// goroutineCount 만큼 고루틴 생성
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		if withPool {
			go workWithPool(&wg)
		} else {
			go workWithoutPool(&wg)
		}
	}

	wg.Done()

	runtime.GC() // 작업 후 GC 실행
	runtime.ReadMemStats(&memAfter)
	duration := time.Since(startTime)

	// 메모리 사용량 및 GC 실행 횟수 차이 계산
	allocDiff := memAfter.Alloc - memBefore.Alloc
	totalAllocDiff := memAfter.TotalAlloc - memBefore.TotalAlloc
	gcDiff := memAfter.NumGC - memBefore.NumGC

	mode := "Without sync.Pool"
	if withPool {
		mode = "With sync.Pool"
	}

	fmt.Printf("==== %s ====\n", mode)
	fmt.Printf("Time elapsed: %v\n", duration)
	fmt.Printf("Allocated bytes diff: %v\n", allocDiff)
	fmt.Printf("Total allocated bytes diff: %v\n", totalAllocDiff)
	fmt.Printf("GC cycles: %v\n", gcDiff)
	fmt.Println()
}

func main() {
	// sync.Pool 미사용 테스트
	runTest(false)
	// sync.Pool 사용 테스트
	runTest(true)
}

/*
결과
	==== Without sync.Pool ====
	Time elapsed: 1.0276ms
	Allocated bytes diff: 991248
	Total allocated bytes diff: 1056768
	GC cycles: 1

	==== With sync.Pool ====
	Time elapsed: 1.5168ms
	Allocated bytes diff: 2345232
	Total allocated bytes diff: 4250856
	GC cycles: 1

분석
Time elapsed : Pool에서 Get,Put 하는 작업 때문에 시간이 더 소요됨 => 즉 객체 생성 비용이, Pool 생성 비용보다 커야 함
Allocated bytes diff / Total allocated bytes diff : sync.Pool에서 재사용 가능한 객체가 없을 때, New로 객체를 생성 하는 데, 이 때문에 추가 적인 할당 발생
*/

// 테스트를 수정해서 webviewer 로직을 가져와야겟네
