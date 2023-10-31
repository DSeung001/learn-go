package main

import "fmt"

func main() {

	slice1 := make([]int, 3, 5)    // len:3, cap:5 슬라이스 생성
	slice2 := append(slice1, 4, 5) // 슬라이스에 값 추가

	// cap으로 capacity값을 확인 가능
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // slice2도 바뀜
	// 슬라이스는 포인터를 사용하므로 같이 바뀜
	fmt.Println("After change sceond element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)
	// cap의 크기를 넘으면 새롭게 슬라이스를 만들기 떄문에 여기는 값을 복사해와서 500을 추가한 값을 새로 만ㄷ름

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
