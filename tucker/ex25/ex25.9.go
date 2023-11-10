package main

import (
	"context"
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func main() {
	wg3.Add(1)

	// context에 값을 추가
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	// 같은 코드에서 아래 코드를 추가해서 값을 할 수 있음
	ctx = context.WithValue(ctx, "number2", 10)
	go square5(ctx)

	cancel()
	wg3.Wait()
}

func square5(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg3.Done()

}
