package main

import "fmt"

// 전체 가로 세로 크기, 배추 밭 크기
var maxX, maxY uint8

func main() {
	// 테스트 케이스 수
	var testCase uint8
	// 배추 위치
	var x, y uint8
	// 배추 수
	var count uint16

	fmt.Scanf("%d\n", &testCase)
	for i := uint8(0); i < testCase; i++ {
		fmt.Scanf("%d %d %d\n", &maxX, &maxY, &count)

		// 2차원 배열 값들 초기화
		// 밭
		field := make([][]uint8, maxX)
		// 방문 여부
		visited := make([][]bool, maxX)
		for i := uint8(0); i < maxX; i++ {
			field[i] = make([]uint8, maxY)
			visited[i] = make([]bool, maxY)
		}
		for i := uint16(0); i < count; i++ {
			fmt.Scanf("%d %d\n", &x, &y)
			field[x][y] = 1
		}

		// 배추흰지렁이 수
		var worm = 0
		for i := uint8(0); i < maxX; i++ {
			for j := uint8(0); j < maxY; j++ {
				if field[i][j] == 1 && !visited[i][j] {
					dfs(field, visited, i, j)
					worm++
				}
			}
		}
		fmt.Println(worm)
	}
}

// go에서 이차원 배열의 할당은 슬라이스 내부적을 배열을 참조하기 때문에 이는 간접적으로 이차원 배열의 포인터를 사용하는것과 같음
// => 포인터로 할당 안받아도 됨
func dfs(filed [][]uint8, visited [][]bool, x, y uint8) {
	// 범위를 벗어난 경우 반환
	if x < 0 || x >= maxX || y < 0 || y >= maxY {
		return
	}
	// 배추가 없거나 방문한 경우 반환
	if filed[x][y] == 0 || visited[x][y] {
		return
	}
	visited[x][y] = true
	// dfs 깊이 우선 탐색으로 모든 방향으로 뻗어나가서 체크를 진행한다
	// 이 부분을 반복문으로 바꿀 수 있지만 가독성이 좋지 않아 패스함
	dfs(filed, visited, x+1, y)
	dfs(filed, visited, x-1, y)
	dfs(filed, visited, x, y+1)
	dfs(filed, visited, x, y-1)
}
