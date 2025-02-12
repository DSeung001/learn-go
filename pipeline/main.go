package main

import (
	"fmt"
	"sync"
	"time"
)

// FanIn 함수: 여러 개의 채널을 하나로 병합하는 역할
func FanIn(sources []<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(sources)) // 모든 채널 개수만큼 WaitGroup 추가

		for _, source := range sources {
			source := source // go 루틴 내부에서 사용하기 위해 변수 복사
			go func() {
				// 채널에서 값을 하나씩 꺼내서 out 채널로 보냄
				for v := range source {
					out <- v
				}
				wg.Done() // 한 채널이 끝나면 Done 호출
			}()
		}

		wg.Wait()  // 모든 고루틴이 종료될 때까지 대기
		close(out) // 결과 채널 닫기
	}()
	return out
}

func main() {
	const number = 10

	// 숫자 생성기: 1부터 n까지의 숫자를 순차적으로 생성하는 채널 반환
	genNumbers := func(n int) <-chan int {
		out := make(chan int)
		go func() {
			for i := 1; i <= n; i++ {
				out <- i
			}
			close(out) // 작업 완료 후 채널 닫기
		}()
		return out
	}

	numChan := genNumbers(number) // 1~10까지의 숫자를 생성하는 채널

	const processDuration = time.Second

	// 숫자를 제곱하는 비동기 처리 함수 (1초 대기 후 값 반환)
	processSquare := func(source <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for v := range source {
				time.Sleep(processDuration) // 1초 대기
				out <- v * v                // 제곱한 값 반환
			}
			close(out)
		}()
		return out
	}

	const fanOutSize = 3
	var fanOut []<-chan int

	// fan-out: 여러 개의 고루틴에서 동시에 숫자를 제곱하도록 설정
	for i := 0; i < fanOutSize; i++ {
		fanOut = append(fanOut, processSquare(numChan))
	}

	// fan-in: 여러 개의 채널에서 나온 값을 하나로 병합
	result := FanIn(fanOut)

	// 결과 출력
	for resultV := range result {
		fmt.Println("result:", resultV)
	}
}
