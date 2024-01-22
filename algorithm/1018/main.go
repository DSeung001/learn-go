package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// 주소 : https://www.acmicpc.net/problem/1018
// 체스판의 시작은 맨위가 흰색, 검은색이고
// 아래 처럼 나와야 함
// 체스판을 8x8로 자르는데 가장 적게 색칠을 해야하는 경우의수를 구하기
// WBWBWBWB
// BWBWBWBW

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var n, m int
	var size int = 8

	fmt.Fscan(reader, &n)
	fmt.Fscanln(reader, &m)
	var board []string
	var results []int

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanln(reader, &str)
		board = append(board, str)
	}

	// 바둑판 찾기
	// 왼쪽 부터 차례대로 시작하는 값으로 인덱스를 가져온 후 반복문
	// 그 반복에서 색칠해야하는 부분 찾는 함수 필요

	// 이렇게 하면 8x8로 자를 때 시작 값을 가질 수 있는데
	for i := 0; i <= n-size; i++ {
		for j := 0; j <= m-size; j++ {

			count := 0
			// 시작값 -1, W면 0, B면 1
			var last rune

			for c := 0; c < size; c++ {
				for r := 0; r < size; r++ {

					target := rune(board[i+c][j+r])
					if r > 0 || c > 0 {
						// 지금 BBB에서 안됨 => 연속된 경우도 생각해야함
						if r == 0 && last == target {
							continue
						} else if !(last == 'W' && target == 'B') && !(last == 'B' && target == 'W') {
							fmt.Println(c, r, string(target), string(last))
							count++

							if last == 'W' {
								last = 'B'
							} else if last == 'B' {
								last = 'W'
							}
							continue
						}
					}
					last = target
				}
			}
			results = append(results, count)
		}
	}

	fmt.Println(slices.Min(results))
}
