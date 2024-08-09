package shortener

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const tableName = "shortener"

type Repository struct {
	*redis.Client
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{client}
}

func (r Repository) Set(ctx context.Context, model *URLInfo) error {
	return r.Client.HSet(ctx, tableName, model.Short, model).Err()
}

func (r Repository) Get(ctx context.Context, key string) (string, error) {
	return r.Client.HGet(ctx, tableName, key).Result()
}

func (r Repository) GetAll(ctx context.Context) (map[string]string, error) {
	return r.Client.HGetAll(ctx, tableName).Result()
}
