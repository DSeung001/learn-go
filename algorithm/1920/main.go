package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func binarySearch(arr []int, elem int) int {
	// 시작과 끝 인덱스를 설정
	start, end := 0, len(arr)-1
	for start <= end {
		// 중간 값을 기준으로 시작과 끝을 변경 및 탐색 반복
		mid := (end + start) / 2
		if arr[mid] > elem {
			end = mid - 1
		} else if arr[mid] < elem {
			start = mid + 1
		} else {
			//arr[mid] == elem 일 경우 return 1
			return 1
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var m int

	// 배열 길이 및 데이터 입력
	// Fscanln 은 줄 바꿈을 기준으로 입력
	// Fscan 은 공백을 기준으로 입력
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
