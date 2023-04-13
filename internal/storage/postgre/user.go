package postgre

import (
	"context"
	"errors"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB     *gorm.DB
	logger logger.RequestLogger
}

func NewUserRepository(DB *gorm.DB, logger logger.RequestLogger) *UserRepository {
	return &UserRepository{DB: DB, logger: logger}
}

func (r *UserRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.DB.WithContext(ctx).Create(&user).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return nil, model.ErrDuplicateKey
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user model.User) error {
	if err := r.DB.WithContext(ctx).Save(user).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return err
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, ID int) error {
	if err := r.DB.WithContext(ctx).Delete(model.User{}, ID).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return err
	}

	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	if err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, ID uint) (*model.User, error) {
	var user model.User

	if err := r.DB.WithContext(ctx).Where("id = ?", ID).First(&user).Error; err != nil {
		r.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return &user, nil
}
