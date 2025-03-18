package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	_struct "grpc_vs_rest/struct"
)

func BenchmarkRestRegisterParallel(b *testing.B) {
	// 동시에 실행되는 고루틴 수 제한 (예: 4)
	b.SetParallelism(4)

	// 단일 연결을 사용하도록 http.Transport 설정
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 200,
			IdleConnTimeout:     90 * time.Second,
			MaxConnsPerHost:     1, // 이 설정으로 단일 연결만 사용하여 임시 포트 고갈 방지
		},
	}

	user := _struct.User{
		LastName:  "Kim",
		FirstName: "Minji",
		Phone:     "010-1234-5678",
		Email:     "minji@example.com",
		Gender:    "F",
		BirthDate: "1990-01-01",
		Username:  "minji90",
	}
	data, err := json.Marshal(user)
	if err != nil {
		b.Fatalf("failed to marshal user: %v", err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := client.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
			if err != nil {
				b.Logf("HTTP POST error: %v", err)
				continue
			}
			resp.Body.Close()
		}
	})
}
