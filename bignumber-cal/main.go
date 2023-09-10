package main

import (
	"fmt"
	"math/big"
)

func main() {
	var number1 = new(big.Int)
	var number2 = new(big.Int)
	var result = new(big.Int)

	// 뒤에 base 파라티머에 경우는 10진수를 의미합니다.
	number1.SetString("18446744073709551", 10)
	number2.SetString("18446744073709551616", 10)

	result.Add(number1, number2)

	fmt.Println(result)
}
