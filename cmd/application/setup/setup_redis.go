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

func (conn *RedisConfig) NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", conn.Host, conn.Port),
		Password:     conn.Password,
		DB:           conn.DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		if conn.Host == "host.docker.internal" {
			client = redis.NewClient(&redis.Options{
				Addr:         fmt.Sprintf("localhost:%s", conn.Port),
				Password:     conn.Password,
				DB:           conn.DB,
				DialTimeout:  5 * time.Second,
				ReadTimeout:  3 * time.Second,
				WriteTimeout: 3 * time.Second,
				PoolSize:     10,
				MinIdleConns: 5,
			})

			if err := client.Ping(ctx).Err(); err != nil {
				return nil, fmt.Errorf("failed to connect to Redis: %v", err)
			}
		} else {
			return nil, fmt.Errorf("failed to connect to Redis: %v", err)
		}
	}

	return client, nil
}
