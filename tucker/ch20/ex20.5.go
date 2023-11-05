package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("v is int: ", t)
	case float64:
		fmt.Println("v is float64: ", t)
	case string:
		fmt.Println("v is string: ", t)
	default:
		fmt.Printf("Not supported type: %t, %v\n", t, t)
	}
}

type Student2 struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello, world!")
	PrintVal(Student2{15})
}
