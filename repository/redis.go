package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(addr string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &Redis{client: client}
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}
