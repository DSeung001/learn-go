// 구조체를 포함한 구조체
package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User  // 필드명을 생략할 경우 바로 접속이 가능해짐, 물론 이름으로도 접속이 가능
	Level int
	Price int
}

func main() {
	user := User{"송하나", "hana", 23, 5}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 3},
		3,
		250, // 여러 줄로 초기화시 쉼표로 마무리 필수, 나중에 값이 추가되었을 때를 위한 세팅
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 : %d VIP 레벨: %d VIP 가격 : %d만 원\n",
		vip.User.Name,
		vip.User.ID,
		vip.User.Age,
		vip.User.Level, // 이름이 겹쳐도 필드명을 적어서 접속 가능
		vip.Price,
	)

	// 구조체 속 구조체는 이름이 겹치지 않으면 바로 접속 가능
	fmt.Printf("VIP 유저: %s ID: %s 나이 : %d VIP 레벨: %d VIP 가격 : %d만 원\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level, // 이름이 겹칠 경우 상단의 구조체에서 먼저 찾음
		vip.Price,
	)
}
