package setup

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisConfig() *RedisConfig {
	host := os.Getenv("REDIS_HOST")
	if host == "localhost" {
		host = "host.docker.internal"
	}

	return &RedisConfig{
		Host:     host,
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
}

func NewRedisClient() (*redis.Client, error) {
	environment := os.Getenv("ENVIRONMENT")
	var client *redis.Client

	if environment == "production" {
		redisURL := os.Getenv("REDIS_URL")
		if redisURL == "" {
			return nil, fmt.Errorf("REDIS_URL not set")
		}
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse REDIS_URL: %v", err)
		}
		client = redis.NewClient(opt)
	} else {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		password := os.Getenv("REDIS_PASSWORD")
		db := 0

		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       db,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return client, nil
}
