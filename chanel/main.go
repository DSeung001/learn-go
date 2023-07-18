package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	// 무한 루프
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			// quit 에 값이 들어오면 인식 (quit 에 값을 뽑을 수 있음)
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
