package main

import "fmt"

func main() {
	// rune은 문자를 표현하는 자료형이지만 한 글자를 표현하기 위해 1~3바이트가 필요합니다
	// 하지만 GO에는 3바이트 정수형이 없기 때문에
	// rune = int32
	var char rune = '한'

	fmt.Printf("%T\n", char) // char 타입 출력, rune = int32
	fmt.Println(char)        // char는 int32라 숫자 출력
	fmt.Printf("%c\n", char) // %c로 문자 출력

	// 한글이 용량을 더 먹음
	str1 := "가다나다라"
	str2 := "abcde"

	// 여기서 len은 문자길이가 아닌 메모리(바이트)의 크기
	fmt.Printf("len(str1) = %d\n", len(str1))
	fmt.Printf("len(str2) = %d\n", len(str2))
}
