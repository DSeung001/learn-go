package main

import (
	"fmt"
	"math"
)

// 1-2. 작은 오차는 무시하기
// 지수부 표현에서 가장 작은 차이는 가장 오른쪽 비트 값하나의 차이이다
// math.Nextafter는 파라미터 중 a가 b보다 작다면 a를 1비트 증가, 크다면 a를 1비트 감소 시켜준다 이를 통해 비교가 가능하다
// 주의) 해당 방법도 무조건 맞는게 아님, 금융 프로그램일 경우 math/big의 float를 사용해서 정밀도를 조정할 수 있음
func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
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
