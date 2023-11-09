package main

import (
	"fmt"
	"net"
)

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		// recover 또한 interface{}
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}

		// 아래와 같이 에러 종류의 정의도 가능
		if r, ok := recover().(net.Error); ok {
			fmt.Println("r is net.Error", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		// 여기서 발생한 panic은 -> g -> f로 역순으로 올라감
		// f에서는 전파된 panic을 복구를 시도함 => 복구시 프로그램이 종료도지 않음
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}
