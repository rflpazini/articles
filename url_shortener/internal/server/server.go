package server

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"url-shortener/internal/monitoring"
	"url-shortener/pkg/api/shortener"
)

type Server struct {
	echo *echo.Echo
	db   *redis.Client
}

func NewServer(db *redis.Client) *Server {
	monitoring.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(monitoring.Middleware)

	monitoring.Route(e)
	shortener.RegisterRoutes(e, db, jsoniter.ConfigCompatibleWithStandardLibrary)

	return &Server{e, db}
}

func (s *Server) Run() error {
	port := ":3001"
	return s.echo.Start(port)
}
