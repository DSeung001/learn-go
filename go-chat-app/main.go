package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"gochatapp/pkg/httpserver"
	"gochatapp/ws"
	"log"
)

// https://levelup.gitconnected.com/create-a-chat-application-in-golang-with-redis-and-reactjs-c75611717f84

func init() {
	// Load the environment file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		fmt.Println("http server is starting on :8080")
		httpserver.StartHTTPServer()
	} else if *server == "websocket" {
		fmt.Println("websocket server is starting on :8081")
		ws.StartWebsocketServer()
	} else {
		fmt.Println("invalid server. Available server : http or websocket")
	}
}
