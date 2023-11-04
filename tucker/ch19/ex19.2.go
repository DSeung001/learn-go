package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	println(a.add(30))
	var b int = 20
	// int에서 myInt로 타입 변환
	fmt.Println(myInt(b).add(50))
}
