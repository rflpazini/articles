package server

import (
	"log"

	"github.com/rflpazini/localstack/internal/config"
	"github.com/rflpazini/localstack/pkg/aws"
)

func InitDependencies() {
	err := aws.CreateTable(config.UsersTable)
	if err != nil {
		log.Printf("create table: %v", err)
	} else {
		log.Println("table created")
	}

	queueURL, err := aws.CreateQueue(config.UsersQueue)
	if err != nil {
		log.Printf("sqs error: %v", err)
	} else {
		config.QueueURL = queueURL
		log.Println("sqs queue created")
	}
}
