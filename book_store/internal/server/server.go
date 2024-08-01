package server

import (
	"book-store/pkg/api/book"
	"book-store/pkg/api/healthcheck"
	"book-store/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	echo *echo.Echo
	db   *pgxpool.Pool
	cfg  *config.Config
}

func NewServer(cfg *config.Config, db *pgxpool.Pool) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	healthcheck.RegisterHealthCheckRoutes(e, db, cfg)
	book.RegisterRoutes(e, db)

	return &Server{e, db, cfg}
}

func (s *Server) Run() error {
	port := ":" + s.cfg.Server.Port
	zap.S().Infof("starting server at port: %s", port)
	return s.echo.Start(port)
}
