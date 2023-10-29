package main

import "fmt"

func main() {
	// 큰 따옴표로 묶으면 특수문자가 작동
	fmt.Print("Hello\tWorld\n")
	// 백쿼트로 묶으면 특수문자가 작동안함
	fmt.Print(`Go is "awesome"!\nGo is Simple and\t 'powerful'`)
	fmt.Print(`
백쿼트는
여러줄도 인식함`)
}
