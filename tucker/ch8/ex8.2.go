package main

import "fmt"

func main() {
	const PI1 float64 = 3.141592653589793238
	var PI2 float64 = 3.141592653589793238

	// PI1 = 4 에러 방지
	PI2 = 4

	fmt.Printf("원주률: %f\n", PI1)
	fmt.Printf("원주률: %f\n", PI2)
}
