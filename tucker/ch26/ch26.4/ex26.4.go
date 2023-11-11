package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {
	lineNo int
	line   string
}

// 파일 내 정보 라인
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.3.exe word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		// 파일 찾기
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("-----------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("-----------------------")
		fmt.Println()
	}
}

// GetFileList : 경로에 매칭와 매칭되는 파일 이름을 가져옴
func GetFileList(pattern string) ([]string, error) {
	filelist := []string{}
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				filelist = append(filelist, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return filelist, nil
}

// FindWordInAllFiles : 경로에 매칭되는 모든 파일에서 단어를 찾아서 반환
func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return findInfos
}

// FindWordInFile : 파일에서 단어를 찾아서 FindInfo 로 반환
func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename: filename, lines: []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file) // 스캐너를 만듦
	for scanner.Scan() {
		line := scanner.Text()
		// 한 줄씩 읽으면 단어 포함 여부 검색
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo: lineNo, line: line})
		}
		lineNo++
	}
	ch <- findInfo
}
