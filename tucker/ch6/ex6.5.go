package main

import "fmt"

func main() {
	var x int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	// int8에서는 127에 +1을 하여 128이되면 최상위 비트가 1이 되므로 오버플로우가 발생해서 오히려 -128이 됨
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, x+1)
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, x+2)
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, x+3)

	var y int8 = -128
	// int8에서는 -128에서 -1을 하면 언더플로우가 발생하여 127이 됨
	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
	fmt.Printf("y\t=%4d, %08b\n", y, y)
	fmt.Printf("y - 1\t=%4d, %08b\n", y, y-1)
}
