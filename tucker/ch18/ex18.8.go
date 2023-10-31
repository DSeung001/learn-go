package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice3 := []int{}
	slice3 = append(slice3, slice1...)

	slice1[1] = 100
	slice2[2] = 200

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
