package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

func main() {
	// CPU 프로파일링 시작 (프로파일링 결과는 "cpu.prof" 파일에 저장)
	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer cpuProfile.Close()

	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	const numGoroutines = 100000 // 10만 개의 goroutine 실행
	var wg sync.WaitGroup

	// 결과를 저장할 버퍼가 있는 채널 생성
	results := make(chan int, numGoroutines)

	// numGoroutines만큼 goroutine 생성하여 동시 작업 수행
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// 간단한 계산: n의 제곱
			result := n * n
			results <- result
		}(i)
	}

	// 모든 goroutine이 완료되면 채널을 닫기 위한 goroutine 실행
	go func() {
		wg.Wait()
		close(results)
	}()

	// 채널에서 결과를 읽어 총합 계산
	total := 0
	for r := range results {
		total += r
	}

	fmt.Println("Total Sum of squares:", total)

	// 프로그램 종료 전 메모리 사용량 출력
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Allocated Memory: %d kb\n", memStats.Alloc/1024)
	fmt.Printf("Total Allocated: %d kb\n", memStats.TotalAlloc/1024)
	fmt.Printf("System Memory: %d kb\n", memStats.Sys/1024)
	fmt.Printf("Number of GC cycles: %d\n", memStats.NumGC)
}
