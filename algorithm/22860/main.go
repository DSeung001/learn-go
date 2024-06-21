package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Node 구조체 정의
type Node struct {
	Parent   string
	Name     string
	IsFolder bool
	Children map[string]*Node
}

// NewNode 함수 정의
func NewNode(parent string, name string, isFolder bool) *Node {
	return &Node{
		Parent:   parent,
		Name:     name,
		IsFolder: isFolder,
		Children: make(map[string]*Node),
	}
}

// AddChild 메서드 정의
func (n *Node) AddChild(newNode *Node) {
	if n.Name == newNode.Parent {
		n.Children[newNode.Name] = newNode
		return
	}

	for _, child := range n.Children {
		child.AddChild(newNode)
	}
}

// CountFiles 메서드 정의
func (n *Node) CountFiles() (int, int) {
	fileSet := make(map[string]struct{})
	var countFilesRecursively func(node *Node)

	countFilesRecursively = func(node *Node) {
		for name, child := range node.Children {
			if child.IsFolder {
				countFilesRecursively(child)
			} else {
				fileSet[name] = struct{}{}
			}
		}
	}
	countFilesRecursively(n)
	return len(fileSet), len(fileSet)
}

// PrintTree 메서드 정의
func (n *Node) PrintTree(indent string) {
	fmt.Printf("%s%s (Folder: %t)\n", indent, n.Name, n.IsFolder)
	for _, child := range n.Children {
		child.PrintTree(indent + "  ")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, m int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

	nodes := NewNode("", "root", true)

	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		parentName := parts[0]
		childName := parts[1]
		isFolder := parts[1] != "file"

		nodes.AddChild(NewNode(parentName, childName, isFolder))
	}

	queries := make([]string, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		queries[i] = scanner.Text()
	}

	// 계층적으로 출력
	nodes.PrintTree("")

	// 쿼리 처리 및 결과 출력
	for _, query := range queries {
		node := findNode(nodes, query)
		if node != nil {
			fileCount, uniqueFileCount := node.CountFiles()
			fmt.Printf("%s %d %d\n", query, fileCount, uniqueFileCount)
		} else {
			fmt.Printf("%s 0 0\n", query)
		}
	}
}

// findNode 함수 정의
func findNode(root *Node, name string) *Node {
	if root.Name == name {
		return root
	}
	for _, child := range root.Children {
		if found := findNode(child, name); found != nil {
			return found
		}
	}
	return nil
}
