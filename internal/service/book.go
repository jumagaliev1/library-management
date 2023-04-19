package service

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage"
)

type BookService struct {
	repo   *storage.Storage
	cfg    config.Config
	logger logger.RequestLogger
}

func NewBookService(repo *storage.Storage, cfg config.Config, logger logger.RequestLogger) *BookService {
	return &BookService{
		repo:   repo,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *BookService) Create(ctx context.Context, book model.Book) (*model.Book, error) {
	s.logger.Logger(ctx).Info("creating book", book)
	return s.repo.Book.Create(ctx, book)
}

func (s *BookService) GetByTitle(ctx context.Context, title string) (*model.Book, error) {
	return s.repo.Book.GetByTitle(ctx, title)
}

func (s *BookService) GetAll(ctx context.Context) ([]model.Book, error) {
	return s.repo.Book.GetAll(ctx)
}
