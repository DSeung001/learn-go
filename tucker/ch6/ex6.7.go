package main

import "fmt"

// 실수 비교에서 오차를 해결하는 방법
// 1-1. 작은 오차는 무시하기
// 해당 방법은 무시할 오차를 정의한다는 것에서 부터 문제가 발생
const epsilon = 0.000001 // 매우 작은 값 = 오차

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.000000000004
	b = 0.000000000002
	c = 0.000000000007

	// %g : 값이 큰 실수 값은 지수 형태로 표현하고 작은 수는 실숫값 그대로 표현
	// %e : 지수 형태로 실수값을 출력(실수만 가능)
	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}
