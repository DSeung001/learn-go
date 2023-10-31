package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0) // 맨뒤 요소추가

	idx := 2 // 추가할 위치
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	slice[idx] = 100
	fmt.Println(slice)

	// [ :idx] 형태는 idex-1 까지의 값까지다
	// 아래 방법으로도 가능하지만 append([]int{100}, slice[idx:]...)를 통해 불필요한 메모리를 소비
	slice2 := append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	fmt.Println(slice2)

	// 다음 방법으로 메모리 소비를 줄일 수 있음
	slice3 := append(slice, 0)
	// slice3의 값을 늘렸으니 아래 방법으로 한칸씩 미룰 수 있다
	copy(slice3[idx+1:], slice3[idx:])
	slice3[idx] = 101
	fmt.Println(slice3)
}
