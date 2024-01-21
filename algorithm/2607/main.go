package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num int
	var origin string

	fmt.Fscanln(reader, &num)
	fmt.Fscanln(reader, &origin)

	var result int

	for i := 1; i < num; i++ {
		var compare string
		fmt.Fscanln(reader, &compare)

		var bigStr, smallStr string
		if len(origin) > len(compare) {
			bigStr, smallStr = origin, compare
		} else {
			bigStr, smallStr = compare, origin
		}

		var check int
		for j := 0; j < len(bigStr); j++ {
			if strings.Contains(smallStr, string(bigStr[j])) {
				smallStr = strings.Replace(smallStr, string(bigStr[j]), "", 1)
			} else {
				check++
			}
		}

		if check <= 1 {
			result++
		}
	}
	fmt.Println(result)
}
