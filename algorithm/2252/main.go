package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 학생 수, 크기 비교 수 입력
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(reader, &n, &m)

	// n+1을 하는 이유는 계산의 편의를 위해 0 인덱스를 사용하지 않기 위함
	// 인접 리스트, 크기 비교를 해서 작은 수(key) -> 큰 수(value)로 되게 저장함
	adjList := make([][]int, n+1)
	// 진입 차수 리스트, 각 노드에 들어오는 진입 차수의 수를 나타냄
	inDegree := make([]int, n+1)

	// 크기 비교들을 입력 받음
	for i := 0; i < m; i++ {
		var s, t int
		fmt.Fscanln(reader, &s, &t)
		adjList[s] = append(adjList[s], t)
		// 큰 값에 진입 차수가 있다는 의미이므로 값을 증가
		inDegree[t]++
	}

	// 큐를 이용한 위상 정렬
	queue := []int{}
	for i := 1; i <= n; i++ {
		// 진입 차수가 0인 것을 큐에 넣음
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 큐가 빌때 까지 정렬)
	for len(queue) > 0 {
		// 큐에서 값을 꺼냄
		current := queue[0]
		queue = queue[1:]
		// 뽑은 값을 그대로 출력
		fmt.Print(current, " ")

		// 큐에서 꺼낸 값과 연결된 노드들의 진입 차수를 1씩 감소
		for _, next := range adjList[current] {
			inDegree[next]--
			// 진입 차수가 0이 되면 큐에 넣음
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
}
