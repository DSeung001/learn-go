package main

import (
	"fmt"
	"sync"
	"time"
)

func square4(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // 1초 간격 시그널
	terminate := time.After(10 * time.Second) // 10초 후 시그널

	for {
		select {
		case <-tick:
			fmt.Printf("Tick.\n")
		case <-terminate:
			fmt.Printf("Terminated.\n")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square4(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Done()
}
