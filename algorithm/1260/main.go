package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Graph struct {
	adjList map[int][]int // 인접 리스트
	visited map[int]bool  // 참여
	order   []int         // 방문 순서
}

// newGraph : Graph 생성
func newGraph() *Graph {
	g := &Graph{}
	g.adjList = make(map[int][]int)
	g.visited = make(map[int]bool)
	return g
}

// addEdge : Graph에 간선 추가
func (g *Graph) addEdge(from, to int) {
	// 양방향 그래프로 이전 값으로 돌아올 수 있게
	g.adjList[from] = append(g.adjList[from], to)
	g.adjList[to] = append(g.adjList[to], from)
}

// 작은 값을 탐색하기 위해 정렬
func (g *Graph) abjSort() {
	for _, v := range g.adjList {
		sort.Ints(v)
	}
}

// print 함수 추가
func (g *Graph) print() {
	for key, value := range g.order {
		if key == len(g.order)-1 {
			fmt.Printf("%d\n", value)
			break
		}
		fmt.Printf("%d ", value)
	}
}

func (g *Graph) dfs(v int) {
	g.visited[v] = true
	g.order = append(g.order, v)

	for _, i := range g.adjList[v] {
		// 인접 리스트 방문(dfs)을 재귀로 반복
		// 방문하지 않은 노드만 방문, bool의 기본 값은 false
		if !g.visited[i] {
			g.dfs(i)
		}
	}
}

// 초기화 함수 추가
func (g *Graph) init() {
	g.order = []int{}
	for i := 0; i < len(g.visited); i++ {
		g.visited[i] = false
	}
}

// bfs 함수 추가
func (g *Graph) bfs(v int) {
	queue := []int{v}
	g.visited[v] = true

	// 큐가 빌 때 까지 반복
	for len(queue) > 0 {
		// 큐에는 깊이가 낮은 값 부터 차례대로 들어가고
		// 깊이가 낮은 걸 방문을 다 해야 다음 깊이로 넘어가짐
		currentVertex := queue[0]
		queue = queue[1:]
		g.order = append(g.order, currentVertex)

		// 반복문 안에서 계속 큐에 값을 추가해줌
		for _, neighbor := range g.adjList[currentVertex] {
			// 방문하지 않은 인접된 노드들을 큐에 추가
			if !g.visited[neighbor] {
				g.visited[neighbor] = true
				queue = append(queue, neighbor)
			}

		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 정점의 수, 간선의 수, 시작 정점
	// 그림에서는 이해를 위해 시작점이 고정된 것 처럼 그렸으나 실제로는 시작점은 없다.
	var n, m, v int
	fmt.Fscanln(reader, &n, &m, &v)

	// Graph 생성
	g := newGraph()

	for i := 0; i < m; i++ {
		var from, to int
		// 간선의 개수만큼 입력 받기
		fmt.Fscanln(reader, &from, &to)
		// 간선은 양방향이므로 양쪽에 모두 추가
		g.addEdge(from, to)
	}
	g.abjSort()

	g.dfs(v)
	g.print()
	g.init()
	g.bfs(v)
	g.print()
}
