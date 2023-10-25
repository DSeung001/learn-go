package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	// 컴퓨터에서는 2진수로 표현하는데
	// 2진수는 소수점 이하의 수를 표현하기 어렵기 때문에 근사치를 사용하게 되므로 오차가 발생할 수 밖에 없음
	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}
