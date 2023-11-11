package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // 스캐너를 통해 한줄 씩 읽기
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
