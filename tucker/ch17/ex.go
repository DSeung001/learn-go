package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Max     = 5000
	Win     = 1000
	Lose    = 100
	Balance = 1000
)

var osStdin = bufio.NewReader(os.Stdin)
var money = Balance
var winText = fmt.Sprintf("You win!, 축하합니다 %d을 획득하셨습니다.", Win)
var loseText = fmt.Sprintf("You lose!, 안타깝네요 %d원을 가져갑니다", Lose)

func inputInt() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		osStdin.ReadString('\n')
	}
	return n, err
}

func main() {

	for {
		fmt.Println("1 ~ 5까지 숫자를 입력하세요.")
		n, err := inputInt()

		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(5)
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if 1 <= n && n <= 5 {
			if n == r {
				fmt.Println(winText)
				money += Win
				if money >= Max {
					fmt.Println("[승리] 축하드립니다. ", money, "을 획득하셨습니다.")
					return
				}
			} else {
				fmt.Println(loseText)
				money -= Lose
				if money <= 0 {
					fmt.Println("[게임오버]")
					return
				}
			}
			fmt.Println("현재 잔액은 ", money, "입니다.")
			fmt.Println("--------------------------")
		} else {
			fmt.Println("1 ~ 5까지만 입력하세요.")
		}
	}

}
