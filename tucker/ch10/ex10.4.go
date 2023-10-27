package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	day := "thursday"

	switch day {
	case "money", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}

	//  switch age := getMyAge(): age 로 숫자 자체를 비교할 수도 있음
	switch age := getMyAge(); {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Tennage")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is ", age)
	}

	// go의 switch에서는 default로 break을 넣어주는데 이걸 빼고 싶으면 fallthrough를 사용하면 됨
	a := 3
	switch a {
	case 1:
		fmt.Println("a == 1")
		break // break를 넣든 말든 달라지지 않음

	case 2:
		fmt.Printf("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough // 보는 사람으로 하여금 혼란을 줄 수 있기에 사용하는걸 추천하진 않음
	case 4:
		fmt.Println("a == 4 ")
	case 5:
		fmt.Println("a == 5")
	default:
		fmt.Println("a > 5")
	}
}
