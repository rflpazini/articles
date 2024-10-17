package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rflpazini/observability/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Run(ctx); err != nil {
			log.Fatalf("Server init error: %v", err)
		}
	}()

	sig := <-sigs
	log.Printf("%v. Shutting down...", sig)

	cancel()

	time.Sleep(2 * time.Second)
}
