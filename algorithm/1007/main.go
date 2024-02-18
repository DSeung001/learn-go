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

// solve : 좌표들로 만들 수 있는 벡터의 합들에서 최소값을 구하는 함수
func solve(points []Point, n int) float64 {
	// 최소값을 구하기 위해 무한대(가장 큰)수로 초기화
	minDistance := math.Inf(1)

	// points의 길이만큼 커지는 2의 거듭제곱, 2^n
	number := (1 << uint(len(points)))

	// 서브넷으로 절반을 나눈다 개념인듯
	for subset := 0; subset < number; subset++ {

		// subset을 2진수로 변환했을 때 1의 개수
		count := bits.OnesCount(uint(subset))

		// count가 n/2와 같다면 => 딱 절반을 나누는 경우의 수
		if count == n/2 {
			sumX, sumY := 0, 0
			for i, p := range points {
				// 좌표 4개면 0000 ~ 1111까지 16가지 경우의 수에서
				// 0011, 0101, 0110, 1001, 1010, 1100 이런식으로
				// 2개씩 묶어서 끝점과 시작점을 나누는 것

				// if) i번째 비트가 1이면 끝점으로 생각하고, else) 0이면 시작점으로 생각
				// 끝점은 더하고
				if subset&(1<<uint(i)) != 0 {
					sumX += p.x
					sumY += p.y
				} else {
					// 시작점은 뺀다
					sumX -= p.x
					sumY -= p.y
				}
			}
			// 두 점사이의 거리를
			// 피타고라스로 계산
			distance := math.Sqrt(float64(sumX*sumX + sumY*sumY))
			// 최소값 갱신
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
