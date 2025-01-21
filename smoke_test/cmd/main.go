package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rflpazini/articles/smoke_test/pkg/api"
)

func main() {
	e := echo.New()

	e.GET("/healthcheck/info", api.HealthCheck)
	e.GET("/catalog", api.GetCatalog)

	e.Logger.Fatal(e.Start(":8080"))
}
