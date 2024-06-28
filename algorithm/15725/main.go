package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	// 입력 문자열을 처리할 때 부호를 유지하도록 -를 +-로 변환
	input = strings.ReplaceAll(input, "-", "+-")
	// +로 나누어 파싱
	// -6x-6 => +-6x+6 => [-6x, 6]
	terms := strings.Split(input, "+")

	coefficient := 0
	for _, term := range terms {
		if len(term) == 0 {
			continue
		}
		// 나눈 항에서 x가 포함되는 지
		if strings.Contains(term, "x") {
			// x의 인덱스를 구하고
			xIndex := strings.Index(term, "x")

			if xIndex == 0 {
				// xIndex가 0이면 계수가 없으니 일차 함수이므로 1로 설정
				coefficient = 1
			} else if xIndex == 1 && term[0] == '-' {
				// xIndex가 1이고 term[0]이 -이면 계수가 0이므로 부호를 유지하고 -1로 설정
				coefficient = -1
			} else {
				// 상수 항의 미분을 적용하지만
				// 차수가 1이므로 계수만 추출하면 됨
				coef, err := strconv.Atoi(term[:xIndex])
				if err != nil {
					fmt.Println("Error parsing coefficient:", err)
					return
				}
				coefficient = coef
			}
		}
	}

	// 상수만 있는 경우 상수 항의 미분으로 0으로 출력
	var result string
	if coefficient != 0 {
		result = fmt.Sprintf("%d", coefficient)
	} else {
		result = "0"
	}

	fmt.Println(result)
}
