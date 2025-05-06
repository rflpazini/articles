package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/rflpazini/articles/translator/internal/api/server"
	"github.com/rflpazini/articles/translator/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	server := api.NewServer(cfg)
	log.Printf("Starting server on port %s", cfg.ServerPort)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server successfully shut down")
}
