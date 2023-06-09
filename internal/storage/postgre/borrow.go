package postgre

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"gorm.io/gorm"
)

type BorrowRepository struct {
	DB     *gorm.DB
	logger logger.RequestLogger
}

func NewBorrowRepository(DB *gorm.DB, logger logger.RequestLogger) *BorrowRepository {
	return &BorrowRepository{
		DB:     DB,
		logger: logger,
	}
}

func (r *BorrowRepository) Create(ctx context.Context, borrow model.Borrow) (*model.Borrow, error) {
	if err := r.DB.WithContext(ctx).Create(&borrow).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return &borrow, nil
}

func (r *BorrowRepository) GetAll(ctx context.Context) ([]model.Borrow, error) {
	var borrows []model.Borrow

	if err := r.DB.WithContext(ctx).Find(borrows).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return borrows, nil
}

func (r *BorrowRepository) GetNoReturned(ctx context.Context) ([]model.Borrow, error) {
	var borrows []model.Borrow
	r.logger.Logger(ctx).Info("Get Not Returned Borrows from database")
	if err := r.DB.WithContext(ctx).Where("returned IS NULL").Find(&borrows).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return borrows, nil
}

func (r *BorrowRepository) GetByTime(ctx context.Context) ([]model.Borrow, error) {
	var borrows []model.Borrow

	if err := r.DB.WithContext(ctx).Where("borrowed >= NOW() - INTERVAL '1 MONTH'").Find(&borrows).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return borrows, nil
}
