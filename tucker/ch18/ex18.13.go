package main

import (
	"fmt"
	"sort"
)

// 구조체 정렬도 가능하려면 Len().Less(), Swap() 메서드를 구현해야 한다.

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18},
	}
	// 인터페이스를 사용해서 정렬기능
	sort.Sort(Students(s))
	fmt.Println(s)
}
