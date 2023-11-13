package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeWebHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // 경로 테스트

	// 핸들러 가져오고
	mux := MakeWebHandler()
	// mux에 요청을 보내고 결과를 받음
	mux.ServeHTTP(res, req)

	// 가져온 결과 테스트
	assert.Equal(http.StatusOK, res.Code)     // code 확인
	data, _ := io.ReadAll(res.Body)           // 데이터를 읽어서 확인
	assert.Equal("Hello World", string(data)) // body 확인
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // 경로 테스트

	// 핸들러 반환
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)   // code 확인
	data, _ := io.ReadAll(res.Body)         // 데이터를 읽어서 확인
	assert.Equal("Hello Bar", string(data)) // body 확인
}
