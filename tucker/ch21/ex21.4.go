package main

import "fmt"

type opFunc func(int, int) int

func getOperator2(op string) opFunc {
	if op == "+" {
		// Go에서는 함수 리터럴, 다른 언어에서는 람다 혹은 익명 함수라고 부름
		return func(i int, i2 int) int {
			return i + i2
		}
	} else if op == "*" {
		return func(i int, i2 int) int {
			return i * i2
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator2("*")

	result := fn(3, 4)
	fmt.Println(result)

	fn2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(fn2(3, 4))

	result2 := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(result2)
}
