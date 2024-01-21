package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func tmp() {
	reader := bufio.NewReader(os.Stdin)
	var num int
	var origin string

	fmt.Fscanln(reader, &num)
	fmt.Fscanln(reader, &origin)

	var result int
	for i := 1; i < num; i++ {
		var compare string
		fmt.Fscanln(reader, &compare)
		if isSame(origin, compare) {
			result++
		}
	}
	fmt.Println(result)
}

func isSame(origin, compare string) bool {
	if len(origin)-len(compare) > 1 || len(compare)-len(origin) > 1 {
		return false
	}

	originSlice := []rune(origin)
	compareSlice := []rune(compare)

	sort.Slice(originSlice, func(i, j int) bool {
		return originSlice[i] < originSlice[j]
	})
	sort.Slice(compareSlice, func(i, j int) bool {
		return compareSlice[i] < compareSlice[j]
	})

	var check int
	if len(origin) < len(compare) {
		originSlice, compareSlice = compareSlice, originSlice
	}

	tmp := make([]rune, 0)
	for i := 0; i < len(originSlice); i++ {
		if i > len(compareSlice)-1 {
			if strings.Contains(string(tmp), string(originSlice[i])) {
				removeChar(tmp, originSlice[i])
				continue
			} else {
				check++
			}
		} else if originSlice[i] != compareSlice[i] {
			tmp = append(tmp, compareSlice[i])
			check++
		}
		if check > 1 {
			return false
		}
	}
	return true
}
