package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Version    = "dev"
	CommitHash = "unknown"
	BuildTime  = "unknown"
)

type HealthResponse struct {
	Status    string `json:"status"`
	ImageRef  string `json:"image_ref,omitempty"`
	EnvName   string `json:"environment,omitempty"`
	Commit    string `json:"commit,omitempty"`
	Version   string `json:"version,omitempty"`
	BuildTime string `json:"build_time,omitempty"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sub := os.Getenv("SUBDOMAIN")
	if sub == "" {
		sub = "demo.127.0.0.1.sslip.io"
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ephemeral preview at http://"+sub+"\n")
	})

	e.GET("/health", func(c echo.Context) error {
		health := HealthResponse{
			Status:    "healthy",
			ImageRef:  os.Getenv("IMAGE_REF"),
			EnvName:   os.Getenv("ENV_NAME"),
			Commit:    CommitHash,
			Version:   Version,
			BuildTime: BuildTime,
		}
		return c.JSON(http.StatusOK, health)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
