package shortener

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"url-shortener/internal/shortener"
)

func RegisterRoutes(e *echo.Echo, client *redis.Client, json jsoniter.API) {
	r := shortener.NewRepository(client)
	s := shortener.Service{Repository: r, Json: json}

	g := e.Group("/v1")

	g.GET("/shortener", s.Get)
	g.POST("/shortener", s.Upsert)
}
