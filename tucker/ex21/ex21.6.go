package main

import "fmt"

// 캡쳐 : 함수 리터럴 내부에서 외부 변수를 참조하는 것
func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		// i는 참조로 됨, 주소 값이 들어간거라 보면 됨
		f[i] = func() {
			fmt.Println(i)
		}
	}
	// 그래서 여기서 실행시 다 3으로 출력
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		// 내부에서 선언했기에 다른 값들이 캡쳐가 됨
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
