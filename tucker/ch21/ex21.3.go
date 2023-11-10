package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// getOperator의 반환 타입이 함수 타입, 합수 타입은 nil도 가능
func getOperator(op string) func(int, int) int { // op에 따른 함수타입을 반환
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입 변수
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)

	// 함수 정의는 다음 처럼 줄여서 사용도 가능 + 매개변수 변수 명은 옵션
	type onFunc func(int, int) int
	operator2 := getOperator("+")
	fmt.Println(operator2(3, 4))
}
