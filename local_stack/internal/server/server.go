package server

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/labstack/echo/v4"
	"github.com/rflpazini/localstack/pkg/api/handlers"
	awsLocal "github.com/rflpazini/localstack/pkg/aws"
)

func Start(cfg aws.Config) {
	e := echo.New()
	awsLocal.InitClients(cfg)

	InitDependencies()

	e.POST("/user", handlers.CreateUser)
	e.GET("/user", handlers.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}
