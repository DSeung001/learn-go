package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// chan 에 신규 값 방지
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Println(cap(c))

	// goroutine 에 chan 사이즈와 채널을 파라미터로
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
