package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := "Internal server error"
		details := ""

		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code

			switch m := he.Message.(type) {
			case string:
				message = m
			case error:
				message = m.Error()
			default:
				message = "Unknown error"
			}
		}

		if code == http.StatusInternalServerError {
			details = "An unexpected error occurred"
		}

		// Log the error
		c.Logger().Error(err)

		// Send the error response
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				c.NoContent(code)
			} else {
				c.JSON(code, ErrorResponse{
					Status:  code,
					Message: message,
					Details: details,
				})
			}
		}
	}
}
