package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	var s = [5]int{1: 10, 3: 30}
	for i := 0; i < 5; i++ {
		fmt.Println(s[i])
	}

	var x = [3]int{10, 20, 30}
	for i := 0; i < 3; i++ {
		fmt.Println(x[i])
	}
}
