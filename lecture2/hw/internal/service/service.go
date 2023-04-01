package service

import (
	"context"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
)

type IRepository interface {
	Create(ctx context.Context, m map[string]interface{}) (*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}

type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
