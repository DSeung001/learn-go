package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	for i := idx + 1; i < len(slice); i++ {
		// 삭제할 인덱스 이후의 요소들을 한 칸씩 앞으로 이동
		slice[i-1] = slice[i]
	}
	// 마지막 요소 지우기
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5, 6}
	slice2 = append(slice2[:idx], slice2[idx+1:]...)
	fmt.Println(slice2)
}
