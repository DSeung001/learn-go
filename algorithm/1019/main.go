package main

import (
	"fmt"
	"strconv"
)

// https://nerogarret.tistory.com/36
func main() {
	var page int
	var page_count int
	counts := make([]int, 10)
	weight := 1

	fmt.Scanln(&page)

	page_count = len(strconv.Itoa(page))

	for i := 0; i < page_count; i++ {

		var replaced, _ = strconv.Atoi(strconv.Itoa(page/10) + "9")
		var remaining = replaced - page

		// 각 자리 수 만큼의 0~9 최대 개수를 더 함
		for index, _ := range counts {
			counts[index] += (replaced/10 + 1) * weight
		}
		// 0~9 중에 초과한 개수를 빼줌
		for i := 10 - remaining; i < 10; i++ {
			counts[i] -= weight
		}

		strNum := strconv.Itoa(page)[:len(strconv.Itoa(page))-1]
		for _, value := range strNum {
			// value에 유니코드가 오므로 -'0'을 해서 int로 변환
			counts[int(value-'0')] -= weight * remaining
		}

		counts[0] -= weight

		page /= 10
		weight *= 10
	}

	for index, count := range counts {
		if index == len(counts)-1 {
			fmt.Printf("%d", count)
			break
		}
		fmt.Printf("%d ", count)
	}
}
