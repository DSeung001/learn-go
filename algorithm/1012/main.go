package main

import "fmt"

func main() {
	var testCase uint8
	var x, y uint8
	var posX, posY uint8
	var result, count uint16

	fmt.Scanf("%d\n", &testCase)
	for i := uint8(0); i < testCase; i++ {
		fmt.Scanf("%d %d %d\n", &x, &y, &count)

		// 2차원 배열 초기화
		grid := make([][]uint8, x)
		for i := uint8(0); i < x; i++ {
			grid[i] = make([]uint8, y)
		}
		for i := uint16(0); i < count; i++ {
			fmt.Scanf("%d %d\n", &posX, &posY)
			grid[posX][posY] = 1
		}

		for i := uint8(0); i < x; i++ {
			for j := uint8(0); j < y; j++ {
				if grid[i][j] == 1 {
					if i+1 < x && grid[i+1][j] == 1 {
						grid[i][j] = 0
					} else if j+1 < y && grid[i][j+1] == 1 {
						grid[i][j] = 0
					}
				}
			}
		}
		result = 0
		for i := uint8(0); i < x; i++ {
			for j := uint8(0); j < y; j++ {
				if grid[i][j] == 1 {
					result++
				}
			}
		}
		print(result)
	}
}
