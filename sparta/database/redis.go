package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var (
	rdClient  *redis.Client
	nDuration = 30 * 24 * 60 * 60 * time.Second
)

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Username: viper.GetString("redis.Username"),
		Password: viper.GetString("redis.Password"),
		DB:       viper.GetInt("redis.DB"),
	})

	_, err := rdClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return rdClient.Set(context.Background(), key, value, nDuration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(context.Background(), key...).Err()
}