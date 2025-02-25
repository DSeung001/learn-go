package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx := context.Background()
	cancelAfter := 10 * time.Second

	ctx, cancel := context.WithCancel(ctx)

	// After the 'cancelAfter' time. call the contexts cancellation function
	time.AfterFunc(cancelAfter, cancel)

	DoWorldDomination(ctx)
	RunLonger()
}

func DoWorldDomination(ctx context.Context) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("Mwah hah ha haaaqahhhhh... Golang at Speed can never be stopped!")
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				err := ctx.Err()
				fmt.Printf("Oh!! '%v'. Those meddling kids!!!\n", err)
				return
			}
		}
	}()
}

func RunLonger() {
	for i := 0; i < 4; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("simulating long running program")
	}
}
