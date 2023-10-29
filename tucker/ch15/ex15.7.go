package main

import "fmt"

func main() {
	// 인덱스를 사용해 바이트 단위 순회
	str1 := "Hello 월드!"

	for i := 0; i < len(str1); i++ {
		// 한글 같은 경우 1~3바이트를 사용하기에 바이트로 순회할 수 없음
		// 한글을 순회 하려면 => []rune으로 변환 또는 range 사용
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", str1[i], str1[i], str1[i])
	}

	// []rune 사용
	arr := []rune(str1)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	}

	// range 사용
	for _, v := range str1 {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", v, v, v)
	}
}
