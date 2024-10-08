package main

import (
	"github.com/rflpazini/localstack/internal/config"
	"github.com/rflpazini/localstack/internal/server"
)

func main() {
	cfg := config.GetAWSConfig()
	server.Start(cfg)
}
