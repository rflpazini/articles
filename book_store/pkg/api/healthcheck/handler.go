package healthcheck

import (
	"net/http"
	"os"
	"runtime/debug"

	"book-store/pkg/config"
	"book-store/pkg/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterHealthCheckRoutes(e *echo.Echo, db *pgxpool.Pool, cfg *config.Config) {
	health := e.Group("/healthcheck")

	health.GET("/info", func(c echo.Context) error {
		return healthCheckHandler(c, db, cfg)
	})
	health.GET("/ping", func(c echo.Context) error {
		return utils.RespondWithJSON(c, http.StatusOK, map[string]string{"status": "ok!"})
	})
}

func healthCheckHandler(c echo.Context, db *pgxpool.Pool, cfg *config.Config) error {
	if err := db.Ping(c.Request().Context()); err != nil {
		return utils.RespondWithJSON(c, http.StatusInternalServerError, map[string]string{"status": "unhealthy", "error": err.Error()})
	}

	var commitHash string
	var goVersion string

	if bi, ok := debug.ReadBuildInfo(); ok {
		goVersion = bi.GoVersion

		for _, kv := range bi.Settings {
			if kv.Key == "vcs.revision" {
				commitHash = kv.Value
				break
			}
		}
	}

	rsp := Response{
		App: App{
			Name:      cfg.Server.AppName,
			Version:   cfg.Server.AppVersion,
			GoVersion: goVersion,
			Codebase: &Codebase{
				CommitHash: commitHash,
				Branch:     os.Getenv("BRANCH_NAME"),
			},
			Environment: &Environment{
				Name:       utils.GetHostName(),
				Region:     os.Getenv("EC2_REGION"),
				InstanceId: os.Getenv("INSTANCE_ID"),
			},
		},
	}

	return utils.RespondWithJSON(c, http.StatusOK, rsp)
}
