package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 데드락 => 서로가 가지고 있는 값을 서로 빼으려고 하는데 서로 주지 않아서 생긱는 교착 현상

var wg2 sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	// 해당 상황에서는 포크, 스푼으로 데드락을 만듦
	// A,B가 존재하는 데 각각 포크, 스푼만 있는데 식사를 하려면 둘 다 있어야하기에 서로가 서로껄 먼저 가지고 싶은데
	// 둘다 주지 않아서 생기는 현상
	wg2.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")
	wg2.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으력 합니다.\n", name)
		first.Lock() // 첫 뮤텍스 획득 시도
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock() // 두번째 뮤텍스 획득 시도
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
}
