package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t, n, m int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		result := 1
		fmt.Fscanln(reader, &n, &m)
		// 서쪽 사이트 만큼이 개수를 가진 수열인데
		// 다리 서로 겹치지 않아야 함 => 앞에서 나왔던 수가 뒤에서 또 나올 수 없음  => 중복이 아님
		for j := 0; j < n; j++ {
			result *= m - j
			result /= j + 1
		}
		fmt.Println(result)
	}
}
