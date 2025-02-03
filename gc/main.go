package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func toMB(b uint64) float64 {
	return float64(b) / (1024 * 1024)
}

func pcInfo() {
	fmt.Printf("PC에서 사용 가능한 CPU 코어: %d, Go에서 사용 가능한 CPU: %d, HeapAlloc(할당된 메모리): %.2fMB, HeapSys(총 메모리): %.2fMB, GC 실행 횟수: %d, 최근 GC 시간(ms): %v\n",
		runtime.NumCPU(), runtime.GOMAXPROCS(0),
		toMB(runtime.MemStats{}.HeapAlloc), toMB(runtime.MemStats{}.HeapSys),
		runtime.MemStats{}.NumGC, debug.GCStats{}.Pause)
}

// gc? 설정은 뭐 아래처럼 되는데 profile도 해봐야할듯?
func main() {
	pcInfo()

	// GC 튜닝: GOGC 값을 50으로 설정
	// 기본 값은 100 = 100%
	// 값이 낮을 수록 GC가 더 자주 발생하여 메모리 사용량은 낮아지지만, CPU 사용량은 증가, 둘은 반비례 관계
	debug.SetGCPercent(15)

	// 메모리 유지할 슬라이스
	var data [][]byte

	// 메모리 할당을 통한 GC 유발 시뮬레이션
	go func() {
		for i := 0; i < 10000; i++ {
			buf := make([]byte, 1024*i) // 1KB씩 메모리 할당
			data = append(data, buf)
		}
	}()

	// 1초마다 GC 통계 출력 (GC 횟수, Heap 메모리 등)
	ticker := time.NewTicker(500 * time.Millisecond)
	for range ticker.C {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)

		fmt.Printf("GC 횟수: %d, HeapAlloc: %.2fMB, 다음 GC 임계값: %.2fMB\n",
			stats.NumGC, toMB(stats.HeapAlloc), toMB(stats.NextGC))

		pcInfo()
		// 예시에서는 10번의 GC 이후 종료합니다.
		if stats.NumGC >= 50 {
			break
		}
	}

	var gcStats debug.GCStats
	debug.ReadGCStats(&gcStats)
	fmt.Printf("최근 GC 시간(ms): %v\n", gcStats.Pause)
}
