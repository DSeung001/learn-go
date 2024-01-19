package ch3

import (
	"sync"
	"testing"
)

// 실행 => go test -bench . -cpu=1

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // 1
		for i := 0; i < b.N; i++ {
			c <- token // 2
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // 3
		for i := 0; i < b.N; i++ {
			<-c // 4
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // 성능 타이머
	close(begin)   // 5
	wg.Wait()
}
