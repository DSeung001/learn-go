package main

import (
	"fmt"
	"math"
	"math/bits"
)

// Point : 좌표 구조체
type Point struct {
	x, y int
}

// solve : 좌표들로 만들 수 있는 벡터의 합의 최소값을 구하는 함수
func solve(points []Point, n int) float64 {
	// 최소값을 구하기 위해 무한대로 초기화
	minDistance := math.Inf(1)

	/*
		(1 << uint(len(points)))는 모든 원소가 포함된 집합을 나타냅니다.
		이 값을 이진수로 표현하면 1이 len(points)개만큼 연속으로 나타납니다.
		따라서 이 값은 모든 부분 집합을 나타내는 비트 마스크로 사용됩니다.
	*/
	// 컴파일해서 프로세스 알아야함
	for subset := 0; subset < (1 << uint(len(points))); subset++ {
		/*
			비트 마스크를 통해 부분집합을 여부를 확인
		*/
		if bits.OnesCount(uint(subset)) == n/2 {
			sumX, sumY := 0, 0
			for i, p := range points {
				// 비트 마스크를 통해 시작 지점 여부를 확인
				// 시작 지점은 더하고 끝 지점은 빼줌
				if subset&(1<<uint(i)) != 0 {
					sumX += p.x
					sumY += p.y
				} else {
					sumX -= p.x
					sumY -= p.y
				}
			}
			distance := math.Sqrt(float64(sumX*sumX + sumY*sumY))
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}

func main() {
	// 테스트 케이스 입력
	var testCase int
	fmt.Scanln(&testCase)

	for i := 0; i < testCase; i++ {
		// 점의 개수 입력
		var number int
		fmt.Scanln(&number)

		// points 좌표 배열 선언
		points := make([]Point, number)
		for j := 0; j < number; j++ {
			// 좌표 입력
			fmt.Scanln(&points[j].x, &points[j].y)
		}

		fmt.Println(solve(points, number))
	}
}
