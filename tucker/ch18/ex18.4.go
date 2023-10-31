package main

import "fmt"

// slice의 내부
/*
	type SliceHeader struct {
		Data uintptr
		Len int // 요소의 개수
		Cap int // 전체길이 capacity 의 약자
	}

	포인터를 사용하기에 배열에 비해서 사용되는 메모리나 속도에 이점이 있음
*/

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	// 5개 중 3개를 사용하지만 나중에 추가될 요소를 위해 2자리 비워둠
	//var slice2 = make([]int, 3, 5)

	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	// slice는 포인터를 통해 값에 접속하므로 데이터를 바구면 바로 적용됨
	fmt.Println(array)
	fmt.Println(slice)
}
