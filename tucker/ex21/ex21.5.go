package main

import "fmt"

// 함수 리터럴 내부 상태
func main() {
	i := 0
	f := func() {
		// 내부에서는 외부 변수의 값 복사가 아닌 인스턴스를 가져옴
		i += 10
	}

	i++
	f()
	fmt.Println(i)
}
