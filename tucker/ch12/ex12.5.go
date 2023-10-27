package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	// a,b가 같은 타입의 같은 메모리(4바이트 * 5 = 20바이트) 물론 이거는 os의 레지스터에 따라 달라지니깐
	// 64비트면 8바이트 => 40바이트가 됨
	b = a
	fmt.Println()
	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
}
