package ch4

import "fmt"

func Test() {
	// 정수형 채널을 반환함
	chanOwner := func() <-chan int {
		// channel의 쓰기 측면을 제한해 다른 채널이 쓰는 걸 제한함
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// 여가서 또 채널을 읽기 전용 뷰로 제한
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving")
	}

	// 즉 채널을 복사해서
	results := chanOwner()
	// 읽기 전용으로 사용함
	consumer(results)

}
