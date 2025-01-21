package api

import (
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

type HealthcheckResponse struct {
	AppName   string `json:"app_name"`
	GoVersion string `json:"go_version"`
	Commit    string `json:"commit,omitempty"`
}

func HealthCheck(c echo.Context) error {
	var commitHash string
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, kv := range bi.Settings {
			if kv.Key == "vcs.revision" {
				commitHash = kv.Value
				break
			}
		}
	}

	rsp := HealthcheckResponse{
		AppName:   "smoke_test_sample",
		GoVersion: runtime.Version(),
		Commit:    commitHash,
	}

	return c.JSON(http.StatusOK, rsp)
}
