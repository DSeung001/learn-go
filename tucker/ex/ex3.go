package main

import (
	"fmt"
	"unsafe"
)

// OS의 시스템 종류에 따라 기본 레지스터의 크기가 다름
// ex) window x32 => 4바이트 window x64 => 8바이트
type test1 struct {
	a int8  // 1바이트
	b int   // 8바이트
	c int16 // 2바이트
	d int   //8바이트
	e int32 // 4바이트
}

// 같은 필드여도 순서를 바꿔주면
type test2 struct {
	a int8  // 1바이트
	c int16 // 2바이트
	e int32 // 4바이트
	b int   // 8바이트
	d int   //8바이트
}

func main() {
	// unsafe.Sizeof는 바이트 크기를 반환합니다.

	// test1이 용량이 더 크게 먹고
	test1 := test1{1, 2, 3, 4, 5}
	fmt.Println("test1은", unsafe.Sizeof(test1), "바이트")

	// test2가 용량을 덜 차지하는 걸 확인할 수 있습니다
	test2 := test2{1, 2, 3, 4, 5}
	fmt.Println("test2는", unsafe.Sizeof(test2), "바이트")
}
