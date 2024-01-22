package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var n, m int
	var size = 8

	fmt.Fscan(reader, &n)
	fmt.Fscanln(reader, &m)
	var board []string

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanln(reader, &str)
		board = append(board, str)
	}
	// 최악의 경우의 수로도 넘을 수 없는 값
	var min = n * m

	// 크기가 8이므로 8*8 체스판을 만들 수 있는 범위까지만 반복
	for i := 0; i <= n-size; i++ {
		for j := 0; j <= m-size; j++ {
			var stone = [2]byte{'W', 'B'}

			// way1: W로 시작하는 체스판
			way1 := 0
			for c := i; c < i+size; c++ {
				for r := j; r < j+size; r++ {
					if board[c][r] != stone[(c+r)%2] {
						way1++
					}
				}
			}

			// way2: B로 시작하는 체스판
			way2 := 0
			for c := i; c < i+size; c++ {
				for r := j; r < j+size; r++ {
					if board[c][r] != stone[(c+r+1)%2] {
						way2++
					}
				}
			}

			// 가장 작은 값 선택
			var results = []int{min, way1, way2}
			min = slices.Min(results)
		}
	}
	fmt.Println(min)
}
