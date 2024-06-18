package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://www.acmicpc.net/problem/22860

type fileInfo struct {
	child  []*fileInfo
	isFile bool
	name   string
}

type fileRequest struct {
	parentFolderName string
	fileName         string
	isFile           uint8
}

func (f *fileInfo) addFile(fileRequest fileRequest) {
	if fileRequest.parentFolderName == f.name {
		if fileRequest.isFile == 1 {
			f.child = append(f.child, &fileInfo{name: fileRequest.fileName, isFile: false})
		} else {
			f.child = append(f.child, &fileInfo{name: fileRequest.fileName, isFile: true})
		}
	} else {
		for _, child := range f.child {
			if !child.isFile {
				child.addFile(fileRequest)
			}
		}
	}
}

func findFile(files *fileInfo, path string) (int, int) {
	var paths = strings.Split(path, "/")

	// 파일 탐색 자체가 잘못됨 => 다시 짜야함
	for _, path := range paths {
		if path == files.name && paths[len(path)-1] == files.name {
			return getNumber(files)
		} else if len(files.child) > 0 {
			return findFile(files, path)
		}
	}
	return 0, 0
}

// 중첩 파일 없애야하고
func getNumber(files *fileInfo) (int, int) {
	fileNum, allNum := 0, 0
	for _, child := range files.child {
		allNum++
		if child.isFile {
			fileNum++
		} else {
			num1, num2 := getNumber(child)
			fileNum += num1
			allNum += num2
		}
	}
	return fileNum, allNum
}

func printFileInfo(fileInfo fileInfo, level int) {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}

	fmt.Printf("level:%d %s%s (File: %t) \n", level, prefix, fileInfo.name,
		fileInfo.isFile)
	for _, child := range fileInfo.child {
		printFileInfo(*child, level+1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var fileNum, folderNum, findNum int
	var fileRequest fileRequest
	var path string
	files := &fileInfo{name: "main"}

	// 파일 수, 폴더 수 입력
	fmt.Fscanln(reader, &fileNum, &folderNum)
	for i := 0; i < fileNum+folderNum; i++ {
		fmt.Fscanln(reader, &fileRequest.parentFolderName, &fileRequest.fileName, &fileRequest.isFile)
		files.addFile(fileRequest)
	}
	printFileInfo(*files, 0)

	fmt.Fscanln(reader, &findNum)
	for i := 0; i < findNum; i++ {
		fmt.Fscanln(reader, &path)
		fmt.Println(findFile(files, path))
	}
}
