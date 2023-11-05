package main

import (
	"html2md.com/utils"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
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

func file2Text(filePath string) string {
	data, err := os.ReadFile(filePath)
	utils.ErrorHandler(err)
	content := string(data)
	return content
}

func byte2html(content []byte, filePath string) {
	err := os.WriteFile(filePath, content, 0)
	utils.ErrorHandler(err)
}

func deleteFile(filePath string) {
	_ = os.Remove(filePath)
}

func main() {
	const (
		resourcePath = "./resource/test.md"
		resultPath   = "./result/test.html"
	)

	md := []byte(file2Text(resourcePath))
	htmlBytes := mdToHTML(md)
	deleteFile(resultPath)
	byte2html(htmlBytes, resultPath)

	fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, htmlBytes)
}
