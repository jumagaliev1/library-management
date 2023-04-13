package postgre

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB     *gorm.DB
	logger logger.RequestLogger
}

func NewBookRepository(DB *gorm.DB, logger logger.RequestLogger) *BookRepository {
	return &BookRepository{
		DB:     DB,
		logger: logger,
	}
}

func (r *BookRepository) Create(ctx context.Context, book model.Book) (*model.Book, error) {
	if err := r.DB.WithContext(ctx).Create(&book).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) GetByTitle(ctx context.Context, title string) (*model.Book, error) {
	var book model.Book

	if err := r.DB.WithContext(ctx).Where("title = ?", title).Find(&book).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetByAuthor(ctx context.Context, author string) (*model.Book, error) {
	var book model.Book

	if err := r.DB.WithContext(ctx).Where("author = ?", author).Find(&book).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetByID(ctx context.Context, ID uint) (*model.Book, error) {
	var book model.Book

	if err := r.DB.WithContext(ctx).Where("id = ?", ID).Find(&book).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}
	return &book, nil
}
