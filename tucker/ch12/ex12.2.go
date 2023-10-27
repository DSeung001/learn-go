package main

import "fmt"

const Y int = 3

// 배열의 크기는 상수여야함

func main() {
	// x := 5
	// a := [x]int{1, 2, 3, 4, 5}  // x는 상수가 아니여서 에러 발생

	b := [Y]int{1, 2, 3}
	var c [10]int

	fmt.Println(b)
	fmt.Println(c)
}
