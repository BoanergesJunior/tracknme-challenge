package repository

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) model.IRepository {
	return &repository{db: db}
}
