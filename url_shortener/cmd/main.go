package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/redis/go-redis/v9"
	"url-shortener/internal/server"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "my-password",
		DB:       0, // use default DB
	})

	s := server.NewServer(client)
	if err := s.Run(); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
}
