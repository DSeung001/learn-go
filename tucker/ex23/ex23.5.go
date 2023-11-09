package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		// 에러의 경우 콜스택으로 불러진 경로가 스택 덕분에 역순으로 출력 됨
		// panic은 interface여서 다 받을 수 있지만 주로 error 타입을 받음
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
