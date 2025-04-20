package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	_struct "grpc_vs_rest/struct"
)

func BenchmarkRestRegisterParallelHTTP1(b *testing.B) {
	// 동시에 실행되는 고루틴 수 제한 (예: 4)
	b.SetParallelism(4)

	// 단일 연결을 사용하도록 http.Transport 설정
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
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
