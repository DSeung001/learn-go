package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go sqaure(&wg, ch)
	ch <- 9
	wg.Wait()
}

func sqaure(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()

	/*
		채널크기가 0인 경우
		데드락 예제
		ch2 := make(chan int)
		ch2 <- 8
		fmt.Println("Never print")*/
}
