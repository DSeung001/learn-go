package main

import "fmt"

func main() {
	// Go에서 지역 변수는 간단한 네이밍 권함
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	if a+b == c {
		fmt.Println("이 조건문은 실행되지 않습니다.")
	}

	// 출력해보면 조건식이 맞을 것 같지만
	fmt.Printf("%f + %f = %f (%v)\n", a, b, a+b, a+b == c)
	// 실제로는 0.1, 0.2, 03는 정확한 0.1, 0.2, 0.3이 아니다. (Go에서 자체 반올림해서 보여주기 때문에 더 햇갈릴 수 있다.)
	fmt.Printf("%0.18f + %0.18f = %0.18f (%v)\n", a, b, c, a+b == c)
}
