package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 문자열 복사는 주소값을 복사하는 것이므로 전체 데이터가 복사될 염려는 안해도 됨
	str1 := "Hello World!"
	str2 := str1

	// go에서 *reflect.StringHeader로의 형변환을 막기 때문에 unsafe.Pointer로 강제로 변환
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
