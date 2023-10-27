package main

import "fmt"

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)          // 재귀 호출
	fmt.Println("After", n) // 재귀 끝나고 호출
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	printNo(3)
	fmt.Println("--------------")
	fmt.Println(factorial(11))
}
