package main

import (
	"container/ring"
	"fmt"
)

// ring은 이미 존재하네
// ring은 오래된 요소는 지워드 되는 조건에 적합 ex) ctrl+z를 통한 취소하기 기능, 리플레이 기능
// List는 Queue, Stack을 만들기 편리할 뿐 직접적인 이름은 아니였으니깐
func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	// 마지막 j = 5가 되는 타이밍에서 r.Next 한 값이 A이다.
	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Prev()
	}
}
