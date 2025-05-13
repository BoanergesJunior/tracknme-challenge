package usecase

import (
	"github.com/BoanergesJunior/tracknme-challenge/internal/http/app/model"
)

type usecase struct {
	repo model.IRepository
}

func New(repo model.IRepository) model.IUsecase {
	return &usecase{
		repo: repo,
	}
}
