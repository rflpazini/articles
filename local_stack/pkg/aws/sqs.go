package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func ListQueues() ([]string, error) {
	result, err := SQSClient.ListQueues(context.Background(), &sqs.ListQueuesInput{})
	if err != nil {
		return nil, err
	}
	return result.QueueUrls, nil
}

func CreateQueue(queueName string) (string, error) {
	result, err := SQSClient.CreateQueue(context.Background(), &sqs.CreateQueueInput{
		QueueName: &queueName,
	})
	if err != nil {
		return "", err
	}
	return *result.QueueUrl, nil
}

func SendMessage(ctx context.Context, queueUrl, messageBody string) error {
	log.Printf("Sending message with body: %s to %s", messageBody, queueUrl)
	_, err := SQSClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: &messageBody,
	})
	return err
}
