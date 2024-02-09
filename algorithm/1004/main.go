package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var testCases int

	// 테스트 케이스 입력
	fmt.Fscanln(reader, &testCases)

	for i := 0; i < testCases; i++ {
		// 시작점, 도착점 입력
		var x1, y1, x2, y2 int
		var result int
		fmt.Fscanln(reader, &x1, &y1, &x2, &y2)

		// 행성 개수
		var n int
		fmt.Fscanln(reader, &n)
		for j := 0; j < n; j++ {
			// 행성 좌표, 반지름 입력
			var cx, cy, r int
			fmt.Fscanln(reader, &cx, &cy, &r)
			if isCenterIn(x1, y1, cx, cy, r) != isCenterIn(x2, y2, cx, cy, r) {
				result++
			}
		}
		fmt.Println(result)
	}
}

// isCenterIn : 행성의 중심이 사각형 안에 있는지 확인
// x, y : 비교 좌표
// cx, cy : 행성 좌표
// r : 행성 반지름
func isCenterIn(x, y, cx, cy, r int) bool {
	// 값의 차가 마이너스여도 제곱이므로 상관 없음
	return (x-cx)*(x-cx)+(y-cy)*(y-cy) < r*r
}
