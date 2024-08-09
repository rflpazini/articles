package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"url-shortener/internal/server"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "my-password",
		DB:       0, // use default DB
	})

	s := server.NewServer(client)
	if err := s.Run(); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
}
