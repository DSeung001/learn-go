package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var m int

	// 배열 길이 및 데이터 입력
	fmt.Fscanln(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fscanln(reader)
	sort.Ints(arr)
	fmt.Fscanln(reader, &m)
	for i := 0; i < m; i++ {
		var elem int
		fmt.Fscan(reader, &elem)
		fmt.Println(binarySearch(arr, elem))

	}
}

func binarySearch(arr []int, elem int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (end + start) / 2
		if arr[mid] > elem {
			end = mid - 1
		} else if arr[mid] < elem {
			start = mid + 1
		} else {
			return 1
		}
	}
	return 0
}
