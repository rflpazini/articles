package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rflpazini/localstack/internal/config"
	"github.com/rflpazini/localstack/pkg/aws"
	"github.com/rflpazini/localstack/pkg/service"
	"github.com/rflpazini/localstack/pkg/service/models"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const (
	userQueueName = "users_queue"
)

func main() {
	ctx := context.Background()
	cfg := config.GetAWSConfig()
	aws.InitClients(cfg)

	queueURL := "http://localhost:4566/000000000000/" + userQueueName

	for {
		messages, err := aws.SQSClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &queueURL,
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     5,
		})
		if err != nil {
			log.Printf("Erro ao receber mensagens: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, msg := range messages.Messages {
			var user models.User
			err := json.Unmarshal([]byte(*msg.Body), &user)
			if err != nil {
				log.Printf("Erro ao desserializar mensagem: %v", err)
				continue
			}

			err = service.CreateUser(ctx, &user)
			if err != nil {
				log.Printf("Create user error: %v", err)
			}

			_, err = aws.SQSClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      &queueURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("Erro ao deletar mensagem: %v", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
