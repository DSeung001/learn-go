package main

import "fmt"

func main() {

	a := 4
	var b float64 = 3.8

	// fmt.Println(a * b) => 에러 발생
	// float64에서 int로 바뀔 때 반내림이 진행 됨
	fmt.Println(a * int(b))

	var c int16 = 3123
	var d int8 = int8(c)

	// int16에서 int8로 바뀌면서 상위 1바이트가 날라가면서 값이 바뀜
	// 00001100 00110011 => 00110011
	fmt.Println(d)

	var e float32 = 1234.567
	var f float32 = 7654.321

	// 실수에는 표현할 수 있는 자릿수가 정해져있어서 수가 제대로 표현되지 않을 수도 있음
	// 원래 값은 9449772.11401 이지만 flot32는 소수부를 6개까지 표현가능하기에 9.449772+e6로 표현됨
	fmt.Println(float32(e * f))
	// 원래도라면 9449772.11401에 4를 곱해서 37799088.456에 가까운 근사치가 나와야 하지만 전혀 다른 값인 37799090이 나옴
	fmt.Println(float32(e * f * 4))

	var g float32 = 123.456784
	// float32가 표현 가능한 수는 8개이므로 소수부가 반올림 처리되서 123.45679로 나옴
	fmt.Println(g)
}
