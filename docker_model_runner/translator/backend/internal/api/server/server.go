package api

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rflpazini/articles/translator/internal/api/handler"
	"github.com/rflpazini/articles/translator/internal/config"
	"github.com/rflpazini/articles/translator/internal/service"

	customMiddleware "github.com/rflpazini/articles/translator/internal/api/middleware"
)

type Server struct {
	echo    *echo.Echo
	cfg     *config.Config
	service *service.TranslatorService
}

func NewServer(cfg *config.Config) *Server {
	translatorService := service.NewTranslatorService(cfg)

	translatorHandler := handler.NewTranslatorHandler(translatorService)

	e := echo.New()

	e.HTTPErrorHandler = customMiddleware.ErrorHandler()

	e.Use(customMiddleware.ConfigureLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.AllowedOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	}))

	api := e.Group("/api")

	api.POST("/translate", translatorHandler.Translate)
	api.GET("/health", translatorHandler.Health)

	return &Server{
		echo:    e,
		cfg:     cfg,
		service: translatorService,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(":" + s.cfg.ServerPort)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
