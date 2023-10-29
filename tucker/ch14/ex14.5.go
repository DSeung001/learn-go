package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}

	// 함수가 끝나면 내부 변수는 가비지 컬랙터가 가져가지만
	// 탈출 검사를 통해 밖에서도 쓰이면 함수의 메모리인 스택 메모리가 아닌 힙 메모리에 저장되서 외부에서 사용 가능
	return &u
}

func main() {
	userPoiner := NewUser("AAA", 23)
	fmt.Println(userPoiner)
}
