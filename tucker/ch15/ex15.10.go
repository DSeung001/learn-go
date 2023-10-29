package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)

	// 문자열 비교는 앞에서 부터 같은 위치에 있는 글자끼리 비교합니다.
	s1 := "BBB"
	s2 := "aaaaAAA"
	s3 := "BBAD"
	s4 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", s1, s2, s1 > s2)
	fmt.Printf("%s < %s : %v\n", s1, s3, s1 < s3)
	fmt.Printf("%s <= %s : %v\n", s1, s4, s1 <= s4)
}
