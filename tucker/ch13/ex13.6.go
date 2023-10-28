package main

import (
	"fmt"
	"unsafe"
)

type User1 struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

// 메모리 정렬로 때문에 2바이트 변수는 2의 배수에 4바이트는 4의 배수 8바이트는 8의 배수로 들어갑니다
// 그렇게 되면 같은 구조체여도 타입의 순서에 따라 메모리 할당량이 다를 수 있습니다.
type User2 struct {
	A int8
	B int
	C int8
	D int
	E int8
}

type User3 struct {
	A int8
	B int8
	C int8
	D int
	E int
}

func main() {
	user1 := User1{22, 23.7}
	fmt.Println(unsafe.Sizeof(user1)) // 12바이트가 아닌 16바이트가 나오는 이유
	// 메모리 정렬을 위해 레지스터 크기의 배수에 따름
	// 현재 내 PC는 64bit 이므로 8바이트 단위로 배수가 되게 메모리에 정렬시킴
	// 12바이트로 되는게 아닌 16바이트의 저장 공간을 할당받은 게 그 이유
	// 이떄 남은 메모리 공간인 4바이트가 "메모리 패딩"으로 들어감
	// 이렇게 하는 이유는 컴퓨터 내부에서 처리하기 용이하기 때문이다

	// 중간 중간에 있는 int8이 패딩을 사용하게 됨으로 40바이트를 차지하게 됩니다.
	user2 := User2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user2))

	// int8인 1바이트가 한군데로 묶이고 거기에 + 5바이트
	// int가 2개니 16바이트
	// 총합 24가 나옵니다
	user3 := User3{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user3))

	// 결론 : 햇갈리면 메모리 값이 작은 필드를 앞에 배치하자
}
