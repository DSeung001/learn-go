package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는 ", avg, "입니다")
}

// 반환 2개 가능
func Divide1(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// 반환할 변수를 미리 지정 가능
func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	PrintAvgScore("김일동", 80, 74, 95)
	PrintAvgScore("송이동", 88, 92, 53)
	PrintAvgScore("박삼동", 78, 73, 78)

	c, success := Divide1(9, 3)
	fmt.Println(c, success)
	d, success := Divide1(9, 0)
	fmt.Println(d, success)

	c, success = Divide2(9, 3)
	fmt.Println(c, success)
	d, success = Divide2(9, 0)
	fmt.Println(d, success)
}
