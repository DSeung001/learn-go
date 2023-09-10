package bignumber

import (
	"strconv"
)

// AddLargeNumbers : 엄청 큰 숫자도 더할 수 있게 해줌
func AddLargeNumbers(a, b string) string {
	// 두 정수를 문자열에서 int 슬라이스로 변환
	number1 := stringToIntSlice(a)
	number2 := stringToIntSlice(b)

	// 항상 number1이 더 큰 수로
	if len(number1) < len(number2) {
		number1, number2 = number2, number1
	}

	// 배열 차이만큼 빈 배열을 만들어서 number2에 앞에 추가하기
	number2 = append(make([]int, len(number1)-len(number2)), number2...)

	carry := 0
	result := make([]int, len(number1))

	// 뒤에서부터 계산
	for i := len(number1) - 1; i >= 0; i-- {
		sum := number1[i] + number2[i] + carry
		result[i] = sum % 10
		carry = sum / 10
	}

	// 결과를 문자열로 변환 + 만약 number1과 number2가 같은 자릿수로 carry 값이 생겼을 때 더하기
	var resultString string
	if carry > 0 {
		resultString = strconv.Itoa(carry)
	}
	for _, digit := range result {
		resultString += strconv.Itoa(digit)
	}

	return resultString
}

// stringToIntSlice : 문자열을 숫자 slice로 변경
func stringToIntSlice(s string) []int {
	var result []int
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		result = append(result, digit)
	}
	return result
}
