package service

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage"
)

type BorrowService struct {
	repo *storage.Storage
	cfg  config.Config
}

func NewBorrowService(repo *storage.Storage, cfg config.Config) *BorrowService {
	return &BorrowService{repo: repo, cfg: cfg}
}

func (s *BorrowService) Create(ctx context.Context, borrow model.Borrow) (*model.Borrow, error) {

	return s.repo.Borrow.Create(ctx, borrow)
}

func (s *BorrowService) GetAll(ctx context.Context) ([]model.Borrow, error) {

	return s.repo.Borrow.GetAll(ctx)
}

func (s *BorrowService) GetNoReturned(ctx context.Context) ([]model.Borrow, error) {

	return s.repo.Borrow.GetNoReturned(ctx)
}


