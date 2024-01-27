package _chan

import (
	"bytes"
	"fmt"
	"os"
)

func Ex1() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintf(&stdoutBuff, "Producer Done.\n")
		// i 값을 미루면 버퍼링 되므로 Producer이 늦게 출력
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
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
