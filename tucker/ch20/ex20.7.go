package main

import "fmt"

type Stringer7 interface {
	String() string
}

type Student7 struct {
	Age int
}

func (s *Student7) String() string {
	return fmt.Sprintf("Stduent age: %d\n", s.Age)
}

func PrintAge(stringer Stringer7) {
	// stringer는 내부에서 *Student7의 인스턴스를 가리키고 있어서 바꿀 수 있음
	// 다시 Student7로 바꾸는 이유는 Age값에 접근하기 위함
	s := stringer.(*Student7)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	// Student7는 Stringer7 인터페이스의 조건을 충족하기에 형변환이 가능
	s := &Student7{15}
	PrintAge(s)
}
