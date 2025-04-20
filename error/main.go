package main

import (
	"fmt"
	"net/http"
)

// gin과 함께 에러를 어떻게 처리 했는지 => https://d2.naver.com/helloworld/6507662?ref=codenary

// 이 예제는 잠재 적인 출력을 잠재 적인 에러와 연결 하는 방법이 핵심
func main() {
	// 고루틴에서 에러가 발생했을 경우 그 상황에 대한 정보를 다 담을 수 있어야 함
	type Result struct {
		Error    error
		Response *http.Response
	}

	/*
		done은 읽기 전용 채널 생성
		반환 타입이 <-chan *http.Respons 이므로 읽기 전용 채널을 반환
	*/
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	errCount := 0
	urls := []string{"https://www.google.com", "https://badhost", "d", "c", "b"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, backing")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
