package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"md2html.com/utils"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type file struct {
	path string
	name string
}

const (
	resourcePath = "resource"
	resultPath   = "result"
)

func mdToHTML(md []byte) []byte {
	// Markdown Parser extensions : 마크다운 파서 설정
	// CommonExtensions : 기본 설정
	// AutoHeadingIDs : 제목에 자동으로 ID를 부여
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	// md 파일 읽기
	doc := p.Parse(md)

	// Html Renderer flags : HTML 렌더링 설정
	// CommonFlags : 기본 설정
	// HrefTargetBlank : 링크를 새 탭에서 열기
	// CompletePage : <html> 태그 추가
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// Convert markdown to HTML : 읽은 마크다운을 HTML로 변환
	return markdown.Render(doc, renderer)
}

// file2Text : 파일을 읽어서 string 으로 반환
func file2Text(filePath string) string {
	data, err := os.ReadFile(filePath)
	utils.ErrorHandler(err)
	content := string(data)
	return content
}

// byte2html : []byte 로 HTML 파일 생성
func byte2html(content []byte, filePath string) {
	err := os.WriteFile(filePath, content, 0)
	// 폴더 생성 코드 추가
	utils.ErrorHandler(err)
}

// deleteFile : 파일 삭제
func deleteFile(filePath string) {
	_ = os.Remove(filePath)
}

// getFileList : 파라미터 경로에서 부터 md 파일을 추출
func getMdFileList(root string) ([]file, error) {
	fileList := []file{}
	// root 경로에서 모든 파일 트리를 탐색하며 함수를 실행
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// 해당 파일이 폴더가 아니고
		if !info.IsDir() {
			// md 확장자인 경우
			mathed, _ := filepath.Match("*.md", filepath.Base(path))
			if mathed {
				fileList = append(fileList, file{path: path, name: info.Name()})
			}
		}
		return nil
	})
	if err != nil {
		return []file{}, err
	}
	return fileList, nil
}

// replaceFromEnd : 뒤에서 부터 old 를 찾아 new 로 변경
func replaceFromEnd(input, old, new string) string {
	lastIndex := strings.LastIndex(input, old)
	if lastIndex == -1 {
		return input
	}
	return input[:lastIndex] + new
}

// createFolderIfNotExists : 폴더가 존재하지 않으면 생성
func createFolderIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 0777 : 모든 권한 부여으로 파일 생성
		// MkdirAll : 재귀함수로 부족한 모든 폴더 생섵
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	if err != nil {
		return err
	}

	return nil
}

// md2html : md 파일을 읽어서 HTML 파일로 변환
func md2html(resourceFile file) error {
	md := []byte(file2Text(resourceFile.path))
	htmlBytes := mdToHTML(md)

	resultFileName := strings.Replace(replaceFromEnd(resourceFile.path, resourceFile.name, ""), resourcePath, resultPath, 1) + replaceFromEnd(resourceFile.name, "md", "html")
	// resultFileName 경로로 폴더가 없을 경우 생성
	err := createFolderIfNotExists(resultFileName)
	if err != nil {
		return err
	}

	deleteFile(resultFileName)
	byte2html(htmlBytes, resultFileName)
	return nil
}

func main() {
	// 해당 경로에서 md 파일 목록을 가져옵니다
	resourcePathList, err := getMdFileList(resourcePath)
	utils.ErrorHandler(err)

	// goroutine 을 기다리기 위해 WaitGroup 생성
	wg := sync.WaitGroup{}
	wg.Add(len(resourcePathList))

	for _, resourceFile := range resourcePathList {
		// 루프 캡처(반복문의 마지막 값으로 go가 실행됨)로 인해 복사본 생성
		resourceFile := resourceFile
		go func() {
			// md 파일을 읽어서 HTML 파일로 변환하는 로직
			err := md2html(resourceFile)
			utils.ErrorHandler(err)
			wg.Done()
		}()
	}

	// goroutine 기다리기
	wg.Wait()
	fmt.Printf("Markdown to HTML")
}
