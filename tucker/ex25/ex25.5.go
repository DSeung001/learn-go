package main

import (
	"fmt"
	"sync"
	"time"
)

func square3(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square3(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// select문을 통해 close(ch)를 대신할 수 있음
	quit <- true
	wg.Done()
}
