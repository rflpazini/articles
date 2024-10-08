package server

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/labstack/echo/v4"
	"github.com/rflpazini/localstack/internal/config"
	"github.com/rflpazini/localstack/pkg/api/handlers"
	awsClients "github.com/rflpazini/localstack/pkg/aws"
)

func Start(cfg aws.Config) {
	e := echo.New()
	awsClients.InitClients(cfg)

	initDependencies()

	e.POST("/user", handlers.CreateUser)
	e.GET("/user", handlers.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func initDependencies() {
	err := awsClients.CreateTable(config.UsersTable)
	if err != nil {
		log.Printf("create table error: %v", err)
	} else {
		log.Println("table created")
	}

	queueURL, err := awsClients.CreateQueue(config.UsersQueue)
	if err != nil {
		log.Printf("create queue error: %v", err)
	} else {
		config.QueueURL = queueURL
		log.Println("sqs queue created")
	}
}
