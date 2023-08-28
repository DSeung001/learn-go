package stringer

import "strconv"

// Revers : 문자열을 역순으로 반환
func Revers(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

// Inspect : 입력값의 길이와 자료형을 반환
func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}
	return inspectNumber(input), "digit"
}

// inspectNumber : 입력된 문자열이 숫자만 일 경우 길이 반환
func inspectNumber(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}
