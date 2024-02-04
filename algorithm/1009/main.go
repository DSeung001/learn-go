package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	results := []int{}
	reader := bufio.NewReader(os.Stdin)
	var re int

	fmt.Fscanln(reader, &re)

	for i := 0; i < re; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		// 1의 자리는 4번 주기로 반복되기에 나누기 연산
		// b는 1 이상인 값이므로 0이 나온다는 건 4라는 의미
		b = b % 4
		if b == 0 {
			b = 4
		}

		// 제곱
		math.Pow(float64(a), float64(b))
		result := int(math.Pow(float64(a), float64(b))) % 10
		if result == 0 {
			// 0일 경우 값을 10으로
			results = append(results, 10)
		} else {
			// 10으로 나눈 나머지 값들
			results = append(results, result)
		}
	}
	for _, value := range results {
		fmt.Println(value)
	}
}
