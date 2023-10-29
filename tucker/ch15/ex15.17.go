package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// string이 불변이기에 아래와 같이 문자열을 수시로 바꾸면 메모리가 많이 소비됨

func main() {
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 := stringheader.Data

	str += " World"
	addr2 := stringheader.Data

	str += " Welcome!"
	addr3 := stringheader.Data

	fmt.Println(str)
	fmt.Printf("addr1 :\t%x\n", addr1)
	fmt.Printf("addr2 :\t%x\n", addr2)
	fmt.Printf("addr3 :\t%x\n", addr3)
}
