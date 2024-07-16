package main

import (
	"os"

	"book-store/internal/logger"
	"book-store/internal/server"
	"book-store/pkg/config"
	"book-store/pkg/database"
	"book-store/pkg/utils"
	"github.com/cristalhq/aconfig"
	"go.uber.org/zap"
)

const (
	pathConfig = "config/"
	fileExt    = ".json"
)

func main() {
	lg, _ := logger.NewZapLogger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			zap.S().Fatal("failed to sync logger", zap.Error(err))
		}
	}(lg)

	undo := zap.ReplaceGlobals(lg)
	defer undo()

	cfg, err := newConfigLoader()
	if err != nil {
		zap.S().Fatal("failed to read config", zap.Error(err))
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		zap.S().Fatalf("failed to connect to database: %v", err)
	}

	s := server.NewServer(cfg, db)
	if err := s.Run(); err != nil {
		zap.S().Fatalf("failed to run server: %v", err)
	}
}

func newConfigLoader() (*config.Config, error) {
	var cfg config.Config

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipEnv:   true,
		SkipFlags: true,
		Files:     []string{pathConfig + utils.GetEnv() + fileExt},
	})
	if err := loader.Load(); err != nil {
		return nil, err
	}

	postgres := os.Getenv("POSTGRES_URI")
	if postgres != "" {
		cfg.DatabaseConfig.URI = postgres
	}

	version := os.Getenv("APP_VERSION")
	if version != "" {
		cfg.Server.AppVersion = version
	}

	return &cfg, nil
}
