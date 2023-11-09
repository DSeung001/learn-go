package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SqureJob struct {
	index int
}

func (j *SqureJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

// 작업분할 방식과 역할 분할 방식으로 뮤텍스 없이도 동시성 프로그래밍을 만들 수 있음
// 영역을 나누는 방식
func main() {
	var jobList []Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SqureJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
