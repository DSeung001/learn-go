package _chan

import (
	"fmt"
)

func Ex1() {
	intStream := make(chan int, 2)
	go func() {
		defer close(intStream)
		defer fmt.Println("Producer Done.")
		// i 값을 미루면 버퍼링 되므로 Producer이 늦게 출력
		for i := 0; i < 5; i++ {
			fmt.Println("Sending:", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("Received: %v.\n", integer)
	}
}

func Ex2() {
	chanOwner := func() <-chan int {
		// intStream := make(chan int, 4)
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for i := 0; i < 5; i++ {
				intStream <- i
			}
		}()
		return intStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
