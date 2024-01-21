package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
	입력
10
ABCD
ABCDD
ABC
DCBAA
DCB
AAAAA
BBBBB
AABBCCDD
AABCDD
A
답 : 4

3
AAACD
AAAAD
DCAAAA

	=> AAAACD
답 2

아래 비교
CD
ACD

결론 : 정렬 기능도 아닌가?
*/

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
		if isLike(origin, compare) {
			result++
		}
	}
	fmt.Println(result)
}

func isLike(origin, compare string) bool {
	if len(origin)-len(compare) > 1 || len(compare)-len(origin) > 1 {
		return false
	}
	longSlice := []rune(origin)
	shortSlice := []rune(compare)

	sort.Slice(longSlice, func(i, j int) bool {
		return longSlice[i] < longSlice[j]
	})
	sort.Slice(shortSlice, func(i, j int) bool {
		return shortSlice[i] < shortSlice[j]
	})

	if len(origin) < len(compare) {
		longSlice, shortSlice = shortSlice, longSlice
	}

	check := 0
	j := 0
	tmp := make([]rune, 0)
	for i := 0; i < len(origin); i++ {
		if i > len(compare)-1 {
			if strings.Contains(string(tmp), string(longSlice[i])) {
				removeChar(tmp, longSlice[i])
				continue
			} else {
				check++
			}
		} else {
			if longSlice[i] != shortSlice[j] {
				tmp = append(tmp, shortSlice[j])
				check++
			}
			j++
			if check > 1 {
				return false
			}
		}
	}
	return true
}

func removeChar(s []rune, target rune) []rune {
	var result []rune

	for _, char := range s {
		if char != target {
			result = append(result, char)
		}
	}

	return result
}
