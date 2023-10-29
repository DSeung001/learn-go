package main

import "fmt"

func main() {
	str := "Hello World"
	// Hello World 문자코드 배열
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))

	str1 := "Hello 월드"
	runes1 := []rune(str1) // []rune 타입으로 변환

	fmt.Printf("len(str1) = %d\n", len(str1))    // string에서의 len은 문자열의 바이트 크기
	fmt.Printf("len(runes) = %d\n", len(runes1)) // runes에서의 len이 글자 개수

}
