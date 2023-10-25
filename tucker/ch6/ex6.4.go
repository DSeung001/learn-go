package main

import "fmt"

func main() {
	var x int8 = 16
	var y int8 = -128
	var z int8 = -1
	var w uint8 = 128

	// 오른쪽 시프트 연산은 2의 승만큼 나누는 것과 동일, 왼족 시프트 연산은 2의 승만큼 곱할때 범위를 초과하면 0이되는 것과 다르게
	// 오른쪽 시프트 연산은 -1를 >> 해도 -1이다.
	// 왼쪽 시프트 연산은 01000000 << 2를 할경우 0이다.
	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", y, y>>2, y>>2)
	fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", z, z>>2, z>>2)
	fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", w, w>>2, w>>2)
}
