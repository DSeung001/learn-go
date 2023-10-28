package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	// 1. 작은 오차 무시하기
	// go에서 실수를 표현할 때 2가지 수가 발생합니다.
	// 이때 두 수의 차이는 마지막 비트 하나 밖에 차이가 발생하지 않으므로 이걸 무시하면 됩니다.
	// Go에서는 편리하게 math 패키지의 Nextafter 함수가 이를 해결해줍니다.
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// math.Nextafter는 첫 번째 인수의 값이 두 번째 인수보다 작으면 1비트 증가 크면 1비트 감소시키고 값을 반환해줍니다.
	// 이걸로 비교 가능합니다.
	fmt.Printf("%0.18f + %0.18f = %0.18f (%v)\n", a, b, c, math.Nextafter(a+b, c) == c)

	// 2. math/big을 사용해서 비교합니다.
	d, _ := new(big.Float).SetString("0.1")
	e, _ := new(big.Float).SetString("0.2")
	f, _ := new(big.Float).SetString("0.3")

	g := new(big.Float).Add(d, e)
	// Cmp는 f가 작으면 -1을 크면 1를 같으면 0을 반환합니다.
	// math/big은 format을 사용시 에러 발생하므로 주의가 필요합니다.
	fmt.Println(d, " + ", e, " = ", f, " (", f.Cmp(g) == 0, ")")

}
