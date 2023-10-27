package main

func main() {
	const C int = 10

	var b int = C * 20
	// C = 20 상수이므로 에러 발생
	// fmt.Println(&C) 상수는 메모리 주솟값에 접근 할 수 없음
}
