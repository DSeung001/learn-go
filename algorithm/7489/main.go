package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var re int

	fmt.Fscanln(reader, &re)

	for re > 0 {
		re--
		var n int
		result := 1
		fmt.Fscanln(reader, &n)

		for i := 1; i <= n; i++ {
			result = getLeftFirstNum(result * i)
		}
		fmt.Println(result % 10)

	}

}

func getLeftFirstNum(n int) int {
	result := n

	for {
		if result%10 == 0 {
			result = result / 10
		} else if result > 1000000 {
			result = result % 1000000
		} else {
			break
		}
	}

	return result
}
