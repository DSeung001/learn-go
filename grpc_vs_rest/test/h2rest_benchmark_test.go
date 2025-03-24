// 이 코드는 golang.org/x/net/http2/h2c 공식 문서를 참고하여 작성한 예시입니다.
package test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/http2"
	_struct "grpc_vs_rest/struct"
)

func BenchmarkRestRegisterParallelHTTP2(b *testing.B) {
	// 동시에 실행되는 고루틴 수 제한 (예: 4)
	b.SetParallelism(4)

	// HTTP/2(h2c) 전용 클라이언트 설정: TLS 없이 HTTP/2 연결 생성
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true, // TLS 없이 HTTP/2 사용 허용
			// DialTLS를 재정의하여 일반 TCP 연결을 반환함으로써 TLS 없이 연결 생성
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
		Timeout: 10 * time.Second,
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
