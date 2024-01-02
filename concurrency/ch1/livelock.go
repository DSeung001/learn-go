package ch1

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func LiveLock() {
	fmt.Println("Livelock Start")

	// 조건 변수 생성, 뮤텍스와 함께 사용
	// 뮤텍스 : 공유 자원에 대한 접근을 제어하는 동기화 기법으로 하나의 고루틴만 접근 가능하게 함
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		// 1 밀리초마다 cadence broadcast => wait를 깨움
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	// cadence lock
	takeStep := func() {
		cadence.L.Lock()
		// cadence wait는 cadence broadcast가 호출될 때까지 대기
		cadence.Wait()
		cadence.L.Unlock()
	}

	// 어떤 사람이 특정 방향으로 움직이도록 시도하는 함수로 성공 여부를 반환
	process := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "%v ", dirName)
		// 방향 값을 1증가 => atomic 패키지는 함수들의 연산이 원자적이다
		atomic.AddInt32(dir, 1)
		// 1밀리세컨드 기다리기
		takeStep()
		// atomic 연산을 했지만 2번 실행되기 때문에 진전을 못함
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()
		// 해당 방향으로 움직일 수 없음을 알게되면 방향 값을 1 감소
		atomic.AddInt32(dir, -1)
		return false
	}

	var resource int32

	start := func(wg *sync.WaitGroup, processName string) {
		// 결과 값을 저장할 변수
		var out bytes.Buffer

		defer func() {
			fmt.Println(out.String())
			wg.Done()
		}()

		// 반복문에 제한을 둠 => 안그러면 계속 돔
		for i := 0; i < 5; i++ {
			if process("loading", &resource, &out) {
				fmt.Fprintf(&out, "Process %v는 정상 실행되었습니다!.", processName)
				return
			}
		}
		fmt.Fprintf(&out, "Process %v에 라이브락이 발생했습니다!.", processName)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go start(&wg, "A")
	go start(&wg, "B")
	wg.Wait()

	fmt.Println("Livelock End")
}
