package config

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

var (
	cfg        aws.Config
	once       sync.Once
	QueueURL   string
	UsersTable = "users"
	UsersQueue = "users_queue"
)

func GetAWSConfig() aws.Config {
	once.Do(func() {
		var err error
		cfg, err = awsConfig.LoadDefaultConfig(context.Background(),
			awsConfig.WithRegion("us-east-1"),
		)
		if err != nil {
			log.Fatalf("error during AWS config: %v", err)
		}
	})
	return cfg
}
