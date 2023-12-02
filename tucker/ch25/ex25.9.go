package main

import (
	"context"
	"fmt"
)

func main() {
	// 채널의 값이 들어오면 이를 리턴하는 함수
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gen이 실행되면서 채널을 만들고 채널 값을 뽑는데, 뽑는 값이 5가 되면 종료
	// 즉 gen의 내부 반복문도 5까지만 실행
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
