package utils

import "github.com/labstack/echo/v4"

func RespondWithJSON(c echo.Context, code int, payload interface{}) error {
	return c.JSON(code, payload)
}

func RespondWithError(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]string{"error": message})
}
