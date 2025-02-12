package main

import "fmt"

func pipeLine1() {
	// 생성자
	// 정수 슬라이스 크기 만큼 동일한 크기의 버퍼링된 정수 채널을 생성
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		// 버퍼링된 채널 => 버퍼 크기 선언
		// 버퍼링된 만큼의 데이터는 미리 저장할 수 있어, 송신이 블로킹되지 않음
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done: // done 채널이 닫히면 종료
					return
				case intStream <- i: // 정수를 intStream에 전송
				}
			}
		}()
		return intStream
	}

	// 곱하기
	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	// 더하기
	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}
	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	// ( n * 2 + 1 ) * 2
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
