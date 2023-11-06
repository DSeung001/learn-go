package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// writeHello 함수 입장에서는 자신이 무슨 기능을 담당할지 모름 => 외부에서 로직을 주입하는 걸 의존성 주입이라함
func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
