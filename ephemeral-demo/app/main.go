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
	Status      string `json:"status"`
	ImageRef    string `json:"image_ref,omitempty"`
	ProjectName string `json:"project_name,omitempty"`
	Commit      string `json:"commit,omitempty"`
	Version     string `json:"version,omitempty"`
	BuildTime   string `json:"build_time,omitempty"`
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sub := os.Getenv("SUBDOMAIN")
	if sub == "" {
		sub = "demo.127.0.0.1.sslip.io"
	}

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ephemeral preview at "+sub+"\n")
	})

	e.GET("/health", func(c echo.Context) error {
		health := HealthResponse{
			Status:      "healthy",
			ImageRef:    os.Getenv("IMAGE_REF"),
			ProjectName: os.Getenv("PROJECT_NAME"),
			Commit:      CommitHash,
			Version:     Version,
			BuildTime:   BuildTime,
		}
		return c.JSON(http.StatusOK, health)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
