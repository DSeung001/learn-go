package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 문자열은 기본적으로 불변이기에 아래처럼 byte로 바꿔서 변경시킨다
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'
	fmt.Println(str)
	fmt.Printf("%s\n", slice)

	// str과 slice는 주소값이 다름
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	// 그래서 주소 값 출력시 다른 걸 확인할 수 있다.
	fmt.Printf("str:\t%x\n", stringheader)
	fmt.Printf("slice:\t%x\n", sliceheader)

}
