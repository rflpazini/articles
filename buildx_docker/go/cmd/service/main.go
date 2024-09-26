package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func usersHandler(c echo.Context) error {
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}

	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	e.GET("/users", usersHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
