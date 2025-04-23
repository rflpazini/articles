package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort     string
	AllowedOrigins []string
	ModelName      string
	ModelEndpoint  string
	LogLevel       string
}

func Load() (*Config, error) {
	godotenv.Load()

	port := getEnv("SERVER_PORT", "3000")
	origins := strings.Split(getEnv("ALLOWED_ORIGINS", "*"), ",")
	modelName := getEnv("MODEL_NAME", "ai/llama3.2:1B-Q8_0")
	modelEndpoint := getEnv("MODEL_ENDPOINT", "http://localhost:12434/engines/llama.cpp/v1/chat/completions")
	logLevel := getEnv("LOG_LEVEL", "info")

	if !strings.HasPrefix(modelEndpoint, "http") {
		if _, err := os.Stat("/var/run/docker.sock"); os.IsNotExist(err) {
			return nil, fmt.Errorf("docker socket not found, required for Docker Model Runner")
		}
	}

	return &Config{
		ServerPort:     port,
		AllowedOrigins: origins,
		ModelName:      modelName,
		ModelEndpoint:  modelEndpoint,
		LogLevel:       logLevel,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
