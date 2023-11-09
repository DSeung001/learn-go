package main

import "fmt"

// error는 사실 아래와 같은 형태의 interface이므로 즉 문자열을 반환하는 Error() 메서드를 가지고 있으면 error로 사용 가능
//type error interface {
//	Error() string
//}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return fmt.Sprintf("암호 길이가 짧습니다. Len:%d RequireLen:%d \n", err.Len, err.RequireLen)
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("myId", "myPw")
	if err != nil {
		// err 인터페이스 반환시 에러 뿐만아니라 필드에 접근할 수 있음
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입했습니다.")
	}
}
