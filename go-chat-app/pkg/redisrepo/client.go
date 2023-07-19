package redisrepo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

// redisClient : redis client 연결 전역변수
var redisClient *redis.Client

// InitialiseRedis : Redis Instance 연결 및 초기화
func InitialiseRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECTION_STRING"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// checking if redis is connected
	pong, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis Connection Failed", err)
	}

	log.Println("Redis Successfully Connected", "Ping", pong)

	redisClient = conn

	return redisClient
}
