package main

import (
	"fmt"
	"time"
)

/*
	예제 주소 : https://hamait.tistory.com/1017
*/

func main() {
	pattern6()
}

// 패턴6 : 파라미터로 채널 값을 넣을 수 있음
var scheduler6 chan string

func consuming6(prompt string) {
	fmt.Println("consuming 호출됨")
	select {
	case scheduler6 <- prompt:
		fmt.Println("이름을 입력받았습니다 : ", <-scheduler6)
	case <-time.After(5 * time.Second):
		fmt.Println("시간이 지났습니다.")
	}
}

func producing6(console chan string) {
	var name string
	fmt.Println("이름 : ")
	fmt.Scanln(&name)
	console <- name
}

func pattern6() {
	console := make(chan string, 1)
	scheduler6 = make(chan string, 1)

	go func() {
		consuming6(<-console)
	}()
	go producing6(console)
	time.Sleep(100 * time.Second)
}

// 패턴4 : 5초 이내에 입력값을 넣지 않으면 만료되는 채널
func pattern4() {
	scheduler := make(chan string)
	go consuming(scheduler)
	go producing(scheduler)

	time.Sleep(100 * time.Second)
}

func consuming(scheduler chan string) {
	select {
	case <-scheduler:
		fmt.Println("이름을 입력받았습니다")
	case <-time.After(5 * time.Second):
		fmt.Println("시간이 만료되었습니다.")
	}
}

func producing(scheduler chan string) {
	var name string
	fmt.Println("이름 : ")
	fmt.Scanln(&name)
	scheduler <- name
}

// 패턴3 : server1, server2 중 가장 먼저되는게 출력
func pattern3() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)

	}
}

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"
}

// 패턴2 : process를 기다림
func pattern2() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1 * time.Second)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}

		scheduling()
	}
}

func process(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "process successful"
}

func scheduling() {
	// do something
}

// 패턴 1 : case 문의 채널 값이 들어올 때 까지 블럭되는 기본문
func pattern1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			c1 <- "one"
		}
	}()

	go func() {
		for {
			time.Sleep(4 * time.Second)
			c1 <- "two"
		}
	}()

	for {
		fmt.Println("start select -----------------")
		// 각 채널에 값이 들어올 때를 핸들링 가능!
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
		fmt.Println("end select -----------------\n\n")
	}
}
