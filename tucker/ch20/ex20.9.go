package main

import "fmt"

type Stringer9 interface {
	String() string
}

type Student9 struct {
}

func (s *Student9) String() string {
	return "Student"
}

type Actor9 struct {
}

func (a *Actor9) String() string {
	return "Actor"
}

func ConvertType(stringer9 Stringer9) {
	// Actor9가 올 경우 Stringer9로 바꿀 수 없음
	// 내부에서 가르키는 값이 *Actor9 타입 인스턴스기 때문
	student := stringer9.(*Student9)
	// *Student9 => Stringer9로 할 순 있지만 반대는 안됨
	fmt.Println(student)
}

func main() {
	actor := &Actor9{}
	ConvertType(actor)
}
