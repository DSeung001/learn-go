package main

import "fmt"

// 2차원 슬라이스
func main() {
	var slices = [][]int{}

	for i := 0; i < 3; i++ {
		var slice = []int{}
		for j := 0; j < 5; j++ {
			slice = append(slice, i+j)
		}
		slices = append(slices, slice)
	}
	fmt.Println(slices)
}
