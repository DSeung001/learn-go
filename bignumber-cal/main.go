package main

import (
	"bignumber-cal.com/bignumber"
	"fmt"
)

func main() {
	// uint64가 아닌 문자열을 사용
	var number1 = "18446744073709551"
	var number2 = "18446744073709551616"

	fmt.Println(bignumber.AddLargeNumbers(number1, number2))
}
