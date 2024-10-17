package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rflpazini/observability/internal/observability"
	"github.com/rflpazini/observability/pkg/api"
)

func Run(ctx context.Context) error {
	observability.InitMeterProvider()
	tp := observability.InitTracer()
	defer func() {
		if err := observability.ShutdownTracerProvider(tp); err != nil {
			log.Fatalf("Error shutting down the TracerProvider: %v", err)
		}
	}()

	e := echo.New()

	e.Use(Middleware())

	api.RegisterRoutes(e)

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down the server...")
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down the server: %v", err)
	}

	return nil
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
