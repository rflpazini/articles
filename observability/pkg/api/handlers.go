package api

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rflpazini/observability/internal/observability"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", HomeHandler)
	e.GET("/metrics", observability.MetricsHandler())
	e.GET("/process", ProcessHandler)
}

func HomeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{Message: "Hello, World!"})
}

func ProcessHandler(c echo.Context) error {
	_, span := otel.Tracer("process-tracer").Start(c.Request().Context(), "ProcessData")
	defer span.End()

	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	if rand.Intn(100) < 20 {
		span.SetStatus(codes.Error, "error")
		response := &Response{
			Message: "Error during processing",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := &Response{
		Message: "completed",
		Data: map[string]interface{}{
			"processed_at": time.Now().Format(time.RFC3339),
			"duration_ms":  rand.Intn(200),
		},
	}

	return c.JSON(http.StatusOK, response)
}
