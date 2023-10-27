package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	// if 초기문은 if 문이 종료되면 변수는 사용 못함
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young ", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age ", age)
	} else if ok {
		fmt.Print("You are beautiful ", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is ", age) age는 소멸되서 사용못함
}
