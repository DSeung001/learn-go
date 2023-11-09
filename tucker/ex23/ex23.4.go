package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 에러 랩핑
// 때론 에러를 감싸서 새로운 에러를 만들어야하는 경우가 있음 => 에러 속에 에러를 숨기는 것
// 에러가 자세히 나옴
// 에러를 감싸서 어떻게 안에 숨겨진 에러를 풀어내는지를 알아보자
func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
		// 에러 감싸기
	}
	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 언어를 읽어서 숫자로 변환하여 반환합니다.
// 변환된 숫자와 읽은 글자 수, 에러를 반환합니다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err)
	}
	return number, len(word), nil
}

func redEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { // 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	redEq("123 3")
	redEq("123 abc")
}
