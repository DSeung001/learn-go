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
	// Fscanln은 줄 바꿈이 있으면 종료
	fmt.Fscanln(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		// Fscan은 공백이 있으면 종료
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fscanln(reader)

	sort.Ints(arr)

	var m int
	fmt.Fscanln(reader, &m)

	for i := 0; i < m; i++ {
		var elem int
		fmt.Fscan(reader, &elem)

		// Search로 2진 탐색으로 시간 단축
		idx := sort.Search(len(arr), func(i int) bool {
			return arr[i] >= elem
		})

		result := idx < n && arr[idx] == elem
		if result {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}
