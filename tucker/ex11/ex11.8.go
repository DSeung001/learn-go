package main

import "fmt"

func main() {
	a := 1
	b := 1
OuterFor: // 레이블 정의
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				// 레이블에 가장 먼저 포함된 for문까지 종료
				break OuterFor
			}
		}
	}

	// 레이블 사용은 편리하나 혼동이 생길 수 있고 버그가 생길 수 있으므로 주의 필요
	// 그러므로 레이블은 꼭 필요한 경우에만 사용
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
