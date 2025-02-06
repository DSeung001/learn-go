package main

import (
	"io"
	"log"
	"pprof.com/fd"
	"sync"
)

// TestApp은 파일 디스크립터 관리를 담당하는 구조체
type TestApp struct {
	files []io.ReadCloser // 열려 있는 파일 리스트
}

// 모든 열린 파일을 닫고, 리스트를 초기화하여 FD 누수를 방지
func (a *TestApp) Close() {
	for _, cl := range a.files {
		_ = cl.Close()
	}
	a.files = a.files[:0] // 파일 디스크립터 절약
}

// 파일을 열고 TestApp에 추가 (pprof 프로파일링 포함)
func (a *TestApp) open(name string) {
	f, err := fd.Open(name) // pprof 프로파일링 포함
	if err != nil {
		log.Printf("File Not Found : %v", err)
		return // 에러 발생 시 추가하지 않음
	}
	a.files = append(a.files, f)
}

// 단일 파일 열기
func (a *TestApp) OpenSingleFile(name string) {
	a.open(name)
}

// 동일한 파일을 10개 연속으로 열기 (테스트용)
func (a *TestApp) OpenTenFiles(name string) {
	for i := 0; i < 10; i++ {
		a.open(name)
	}
}

// 파일을 100개 동시(고루틴)로 열기 (멀티스레드 테스트)
func (a *TestApp) Open100FilesConcurrently(name string) {
	wg := sync.WaitGroup{}
	wg.Add(10) // 10개의 고루틴 실행
	for i := 0; i < 10; i++ {
		go func() {
			a.OpenTenFiles(name) // 각 고루틴이 10개씩 파일을 엶
			wg.Done()
		}()
	}
	wg.Wait() // 모든 고루틴이 끝날 때까지 대기
}

func main() {
	a := &TestApp{}
	defer a.Close() // 실행 종료 시 모든 열린 파일 닫기

	testFilePath := "./test.txt"

	// 프로파일링을 위해 반복적인 파일 열기 및 닫기
	for i := 0; i < 10; i++ {
		a.OpenTenFiles(testFilePath)
		a.Close() // 닫기 호출하여 FD 누수 방지
	}

	// pprof 분석을 위해 실제로 열릴 파일
	f, _ := fd.Open(testFilePath)
	a.files = append(a.files, f)

	// 다양한 파일 열기 시나리오 실행
	a.OpenSingleFile(testFilePath)
	a.OpenTenFiles(testFilePath)
	a.Open100FilesConcurrently(testFilePath)

	// 현재 파일 디스크립터 사용 내역을 pprof 파일로 저장
	if err := fd.Write("fd.pprof"); err != nil {
		log.Fatal(err)
	}
}
