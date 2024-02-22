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

	// 입력 값을 문자열로 바꿔서 자릿수 만큼 반복하게 수정
	page_count = len(strconv.Itoa(page))

	for i := 0; i < page_count; i++ {
		// 현재 page 값을 기준으로 1의 자리수를 9로 만들고
		var replaced, _ = strconv.Atoi(strconv.Itoa(page/10) + "9")
		// 차를 구함
		var remaining = replaced - page

		// 일의 자리수 0~9가 나오는 최대 반복 회수 저장
		for index, _ := range counts {
			counts[index] += (replaced/10 + 1) * weight
		}

		// 일의 자리수의 반복수가 초과된 만큼 빼기
		for i := 10 - remaining; i < 10; i++ {
			counts[i] -= weight
		}

		// 일의 자리수를 제외한 나머지 자릿수에서 일의 자릿수가 반복되지
		// 않아서 생기는 수 만큼 뺴기
		strNum := strconv.Itoa(page)[:len(strconv.Itoa(page))-1]
		for _, value := range strNum {
			counts[int(value-'0')] -= weight * remaining
		}

		// 수는 1부터 시작하므로
		// 0이 자릿수만큼 반복되는 경우는 없음 (십의 자리, 백의 자리 이상도 마찬가지)
		counts[0] -= weight

		// 페이지를 줄이고
		page /= 10
		// 자릿수를 증가 시킴
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
