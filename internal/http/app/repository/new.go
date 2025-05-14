package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) model.IRepository {
	return &repository{
		db:    db,
		redis: redis,
	}
}
