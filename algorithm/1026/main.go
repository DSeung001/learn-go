package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	A := make([]int, n)
	B := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&B[i])
	}

	sort.Ints(A)
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	S := 0
	for i := 0; i < n; i++ {
		S += A[i] * B[i]
	}
	
	fmt.Println(S)
}
