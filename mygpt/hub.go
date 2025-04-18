package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
		case message := <-h.broadcast:
			for conn := range h.clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					conn.Close()
					delete(h.clients, conn)
				}
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 1) HTTP 연결을 WebSocket으로 업그레이드
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade 실패:", err)
		return
	}
	// 2) 허브에 신규 연결 등록
	hub.register <- conn

	// 3) 별도 고루틴에서 메시지 수신 루프 실행
	go func() {
		defer func() {
			// 연결 해제 시 허브에 알림
			hub.unregister <- conn
		}()
		for {
			// 클라이언트 메시지 읽기
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			// 4) handleOllama에 conn, hub, 메시지를 전달
			go handleOllama(conn, hub, msg)
		}
	}()
}
