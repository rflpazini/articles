package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rflpazini/localstack/internal/config"
	awsClient "github.com/rflpazini/localstack/pkg/aws"
	"github.com/rflpazini/localstack/pkg/service"
	"github.com/rflpazini/localstack/pkg/service/models"
)

func GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.QueryParam("email")
	if email == "" {
		users, err := service.GetAllUsers()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}

	user, err := service.GetUserByEmail(ctx, email)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	message, err := json.Marshal(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = awsClient.SendMessage(ctx, config.QueueURL, string(message))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}
