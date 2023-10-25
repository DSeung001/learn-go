package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 만약 숫자를 하나 입력 받는데 HELLO가 입력으로 온다면 H에서 숫자가 아니므로 에러가 발생
	// 그러면 버퍼에 남은 ELLO는 계속있길래 이걸 지워져야 함, 안그러면 다시 받게 됨
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력 스트림 지우기, 줄바꿈이 나올때까지 입력 버퍼를 읽음
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
