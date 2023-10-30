package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func inputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	numbers := []int{}
	cnt := 0

	// 시간을 기준으로 난수 생성
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)

	for {
		fmt.Printf("숫자값을 입력하세요: ")
		x, err := inputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			cnt++
			numbers = append(numbers, x)
			if x > n {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if x < n {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println(cnt, "번만의 숫자를 맞췄습니다.")
				fmt.Println("입력하신 숫자 : ", numbers)
				break
			}
		}
	}

}
