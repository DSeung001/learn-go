package main

import "fmt"

// 슬라이싱은 자르는 동작 [:] 이고 슬라이스는 결과물
// 시작인덱스 부터 끝인덱스 -1 까지 포함함
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100
	fmt.Println("After change second element")
	fmt.Println("array:", array)
	// slice의 cap은 array의 cap에서 시작 인덱스를 뺀 만큼 할당 받음
	fmt.Println("slice:", slice, len(slice), cap(slice))

	// cap이 4여서 새롭게 메모리를 할당하지 않고 그대로 값을 추가함 => 기존에 배열의 메모리를 침해
	// => slice는 값이 추가되고 array는 같이 바뀜 array[2] = 500 꼴이 됨
	slice = append(slice, 500)
	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice1 := []int{1, 2, 3, 4, 5}
	// slice2는 인덱스 1부터 2까지 가져오기에 [2,3]이 되고 최대 인덱스가 4이므로
	// 시작 인덱스 1를 빼 cap을 3가지게 됩니다.
	slice2 := slice1[1:3:4]
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

}
