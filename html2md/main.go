package main

import (
	"html2md.com/utils"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
)

var printAst = false

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	if printAst {
		fmt.Print("--- AST tree:\n")
		ast.Print(os.Stdout, doc)
		fmt.Print("\n")
	}

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

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
