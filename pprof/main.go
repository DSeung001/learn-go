package main

import (
	"io"
	"log"
	"pprof.com/fd"
	"sync"
)

type TestApp struct {
	files []io.ReadCloser
}

func (a *TestApp) Close() {
	for _, cl := range a.files {
		_ = cl.Close()
	}
	a.files = a.files[:0] // 파일 디스크립터 절약
}

// fd.Open 함수를 사용하여 파일 개방, 이 함수는 파일을 오픈 함과 프로파일링 측정을 같이 함
func (a *TestApp) open(name string) {
	f, _ := fd.Open(name)
	a.files = append(a.files, f)
}

func (a *TestApp) OpenSingleFile(name string) {
	a.open(name)
}

func (a *TestApp) OpenTenFiles(name string) {
	for i := 0; i < 10; i++ {
		a.open(name)
	}
}

func (a *TestApp) Open100FilesConcurrently(name string) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			a.OpenTenFiles(name)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	a := &TestApp{}
	defer a.Close()

	// No matter how many files we opened in the past...
	// 프로파일링 작업을 점검하기 위해 먼저 10개의 파일을 열고 닫고를 10번 반복하는데, 테스트 목적이므로 임의의 파일을 지정
	for i := 0; i < 10; i++ {
		a.OpenTenFiles("/dev/null")
		a.Close()
	}

	// ...after the last Close, only files below will be used in profile
	testFilePath := "./test.txt"
	f, _ := fd.Open(testFilePath)
	a.files = append(a.files, f)

	a.OpenSingleFile(testFilePath)
	a.OpenTenFiles(testFilePath)
	a.Open100FilesConcurrently(testFilePath)

	if err := fd.Write("fd.pprof"); err != nil {
		log.Fatal(err)
	}
}
