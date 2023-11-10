package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
	}
	
	// defer의 호출은 아래부터 위로, 즉 역순
	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello, world!를 씁니다.")
	fmt.Fprintln(f, "Hello, world!")
}
