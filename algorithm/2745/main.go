package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	var n int
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &n)
	var result uint32

	for i := 0; i < len(s); i++ {
		l := len(s) - i - 1
		if s[i] >= 65 {
			result += uint32(s[i]-55) * pow(n, l)
		} else {
			result += uint32(s[i]-48) * pow(n, l)
		}

	}

	fmt.Println(result)
}

func pow(a, b int) uint32 {
	if b == 0 || a == 1 {
		return 1
	}
	if b == 1 {
		return uint32(a)
	}
	n := a
	for i := 2; i <= b; i++ {
		a *= n
	}
	return uint32(a)
}
