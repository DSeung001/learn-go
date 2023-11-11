package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/*
1. csv 파일 읽기 (파일명: resource/sample.csv)
2. 각 라인별로 수스자가 아닌 값이 있으면 해당 라인은 출력하지 않음
3. 출력할 때 다음 항목별로 계산을 함
	- 최소값
 	- 최대값
	- 합계
	- 평균
	- 표준 편차
	- 중간 값
	- 중간 값
4. 출력은 다음과 같이 함
	최소값, 최대값, 합계, 평균, 표준 편차, 중간 값, 중간 값
5. 숫자를 출력 한 후 출력되지 않은 라인을 차레대로 출력함
*/

// csv 파일 읽기
func readCsvFile(fileName string) *[][]string {
	// 파일 열기
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("파일 열기 오류 발생")
		fmt.Println(err)
	}

	// 파일 읽기
	rds := csv.NewReader(bufio.NewReader(file))
	rds.FieldsPerRecord = -1
	rows, err := rds.ReadAll()

	if err != nil {
		fmt.Println("파일 읽기 오류 발생")
		fmt.Println(err)
	}

	return &rows
}

func main() {
	rows := readCsvFile("resource/sample.csv")

	var numbers [][]int
	var strings [][]string

	for _, row := range *rows {

		c := true
		number := make([]int, 0, 10)

		for _, col := range row {
			v, err := strconv.Atoi(col)
			if err != nil {
				c = false
				break
			} else {
				number = append(number, v)
			}
		}

		if c {
			numbers = append(numbers, number)
		} else {
			strings = append(strings, row)
		}
	}

	//fmt.Println(numbers)
	// 왜 다른 값도 걸리지? => 한번 체크해보자
	fmt.Println(strings)
}
