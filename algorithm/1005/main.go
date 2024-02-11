package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 테스트 케이스만큼 반복
	reader := bufio.NewReader(os.Stdin)
	var testCases int
	fmt.Fscanln(reader, &testCases)
	for i := 0; i < testCases; i++ {
		acmCraft(reader)
	}
}

func acmCraft(reader *bufio.Reader) {
	// 건물 수 n, 건물 순서 k, 목표 건물 번호 w 입력
	var n, k, w int

	// 건물 수, 순서 입력
	fmt.Fscanln(reader, &n, &k)

	// n+1을 하는 이유는 계산의 편의를 위해 0 인덱스를 사용하지 않기 위함
	// 인접 리스트, 크기 비교를 해서 작은 수(key) -> 큰 수(value)로 되게 저장할 예정
	adjList := make([][]int, n+1)
	// 건물 당 건설에 걸리는 시간
	buildTime := make([]int, n+1)
	// 진입 차수
	inDegree := make([]int, n+1)
	// 걸리는 시간
	buildTotalTime := make([]int, n+1)

	// 건물 당 건설에 걸리는 시간 입력
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &buildTime[i])
	}
	fmt.Fscanln(reader)

	// 건물 순서 입력
	for i := 0; i < k; i++ {
		var prev, next int
		fmt.Fscanln(reader, &prev, &next)
		// 건물 번호로 작은 수 -> 큰 수로 저장
		adjList[prev] = append(adjList[prev], next)
		// 큰 수의 진입 차수 증가
		inDegree[next]++
	}

	// 목표 건물 번호 입력
	fmt.Fscanln(reader, &w)

	// 진입 차수가 0인 것을 큐에 넣음
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
			buildTotalTime[i] = buildTime[i]
		}
	}

	// 목표 건물을 바로 지을 수 있는 경우 반환
	if inDegree[w] == 0 {
		fmt.Println(buildTime[w])
		return

	}

	// 큐가 빌때 까지 정렬
	for len(queue) > 0 {
		// 큐에서 값을 꺼냄
		current := queue[0]
		queue = queue[1:]

		// 큐에서 꺼낸 값과 연결된 노드들의 진입 차수를 1씩 감소
		for _, next := range adjList[current] {
			// 각 노드별로 가장 오래 걸리는 시간을 저장
			if buildTotalTime[next] < buildTotalTime[current]+buildTime[next] {
				buildTotalTime[next] = buildTotalTime[current] + buildTime[next]
			}
			// 진입 차수가 0이 되면 큐에 넣음
			inDegree[next]--
			if inDegree[next] == 0 {
				// 만약 목표 건물이 진입 차수가 0이 되면 그 건물의 건설 시간을 출력하고 종료
				if next == w {
					fmt.Println(buildTotalTime[next])
					return
				}
				queue = append(queue, next)
			}
		}
	}

}
