package main

import "fmt"

// 타입 없는 상수는 변수에 복사될 때 타입이 정해지기에 여러가지 타입으로 사용되는 상수에서 사용하기 용이
const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	// 상수는 리터럴로 변환되기에 아래 문구는 컴파일 단계에서 다음과 같이 변환 : var a int = 314
	// 상수는 리터럴로 변환되기에 메모리 동적할당이 없고 그래서 &을 통해 메모리 주소값을 알아낼 수 없습니다.
	var a int = PI * 100
	// var b int = FloatPI * 100 => 타입이 달라서 연산 x

	fmt.Println(a)
	// fmt.Println(b)
}
