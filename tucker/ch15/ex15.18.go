// Builder를 통하면 메모리 낭비를 없애고 문자열을 더 할 수 있음
// 문자열에서 불변 원칙을 지키려는 이유는 문자열이 바뀌었을 때 전체다 변경되는 사태를 마아 버그를 줄이기 위함입니다

package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {

		// 아래 코드는 매번 신규 문자열을 메모리에 할당 함
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
