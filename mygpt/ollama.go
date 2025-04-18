package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ollama/ollama/api"
)

// ChatRequest/ChatResponse 구조체는 ollama/api 패키지의 정의를 그대로 사용합니다.

// handleOllama는 클라이언트 메시지를 받아 Ollama 스트리밍 응답을 브로드캐스트합니다.
func handleOllama(conn *websocket.Conn, hub *Hub, userMsg []byte) {
	// 1) URL 파싱 및 클라이언트 초기화
	baseURL, err := url.Parse("http://localhost:11434")
	if err != nil {
		log.Fatalf("URL 파싱 실패: %v", err)
	}
	client := api.NewClient(baseURL, http.DefaultClient)

	// 2) 타임아웃 컨텍스트 생성
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 3) ChatRequest를 포인터로 생성
	req := &api.ChatRequest{
		Model: "llama3.2",
		Messages: []api.Message{
			{Role: "user", Content: string(userMsg)},
		},
	}

	// 4) Chat 호출: 콜백(fn)으로 스트리밍 응답 처리
	if err := client.Chat(ctx, req, func(resp api.ChatResponse) error {
		// resp.Message.Content에 담긴 토큰(혹은 청크)을 브로드캐스트
		hub.broadcast <- []byte(resp.Message.Content)
		return nil
	}); err != nil {
		log.Println("Ollama.Chat 오류:", err)
	}

	// 5) 응답 완료 후 완료 신호 전송
	hub.broadcast <- []byte("\n\n__MESSAGE_COMPLETE__")
}
