package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "non_existent_file.txt"

	_, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일 열기 오류 : ", err)
		return
	}
}
