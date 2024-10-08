package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/rflpazini/localstack/internal/config"
	awsClient "github.com/rflpazini/localstack/pkg/aws"
	"github.com/rflpazini/localstack/pkg/service/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	existingUser, err := GetUserByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email is already in use by another user")
	} else if err != nil && err.Error() != "user not found" {
		return fmt.Errorf("failed to verify if email is already in use: %w", err)
	}

	user.ID = uuid.NewString()

	item := map[string]types.AttributeValue{
		"ID":       &types.AttributeValueMemberS{Value: user.ID},
		"Name":     &types.AttributeValueMemberS{Value: user.Name},
		"Email":    &types.AttributeValueMemberS{Value: user.Email},
		"Password": &types.AttributeValueMemberS{Value: user.Password},
		"Address":  &types.AttributeValueMemberS{Value: user.Address},
		"Phone":    &types.AttributeValueMemberS{Value: user.Phone},
	}

	_, err = awsClient.DynamoDBClient.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(config.UsersTable),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	result, err := awsClient.DynamoDBClient.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(config.UsersTable),
		FilterExpression: aws.String("Email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}

	if len(result.Items) == 0 {
		return nil, errors.New("user not found")
	}

	item := result.Items[0]
	user := &models.User{
		ID:      item["ID"].(*types.AttributeValueMemberS).Value,
		Name:    item["Name"].(*types.AttributeValueMemberS).Value,
		Email:   item["Email"].(*types.AttributeValueMemberS).Value,
		Address: item["Address"].(*types.AttributeValueMemberS).Value,
		Phone:   item["Phone"].(*types.AttributeValueMemberS).Value,
	}

	return user, nil
}

func GetAllUsers() ([]*models.User, error) {
	result, err := awsClient.DynamoDBClient.Scan(context.Background(), &dynamodb.ScanInput{
		TableName: aws.String(config.UsersTable),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all users: %w", err)
	}

	if len(result.Items) == 0 {
		return nil, errors.New("no users found")
	}

	users := make([]*models.User, 0)

	for _, item := range result.Items {
		user := &models.User{
			ID:      item["ID"].(*types.AttributeValueMemberS).Value,
			Name:    item["Name"].(*types.AttributeValueMemberS).Value,
			Email:   item["Email"].(*types.AttributeValueMemberS).Value,
			Address: item["Address"].(*types.AttributeValueMemberS).Value,
			Phone:   item["Phone"].(*types.AttributeValueMemberS).Value,
		}
		users = append(users, user)
	}

	return users, nil
}
