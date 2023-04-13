package service

import (
	"context"
	"errors"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage"
)

type TransactionService struct {
	repo   *storage.Storage
	user   *UserService
	logger logger.RequestLogger
}

func NewTransactionService(repo *storage.Storage, user *UserService, logger logger.RequestLogger) *TransactionService {
	return &TransactionService{repo: repo, user: user, logger: logger}
}

func (s *TransactionService) Create(ctx context.Context, transaction model.Transaction) (*model.Transaction, error) {
	user, err := s.user.GetUserFromRequest(ctx)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return nil, err
	}

	transaction.UserID = user.ID

	if !s.checkBalance(*user, transaction.Amount) {
		s.logger.Logger(ctx).Error("not enough money")
		return nil, errors.New("not enough money")
	}

	s.updateBalance(user, -transaction.Amount)
	if err := s.repo.User.Update(ctx, *user); err != nil {
		s.logger.Logger(ctx).Error(err)
		return nil, err
	}

	return s.repo.Transaction.Create(ctx, transaction)
}

func (s *TransactionService) GetByID(ctx context.Context, ID uint) (*model.Transaction, error) {

	return s.repo.Transaction.GetByID(ctx, ID)
}

func (s *TransactionService) Cancel(ctx context.Context, transactionID uint) error {
	transaction, err := s.GetByID(ctx, transactionID)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}

	user, err := s.user.GetUserFromRequest(ctx)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}
	transaction.UserID = user.ID

	s.updateBalance(user, transaction.Amount)

	if err := s.repo.User.Update(ctx, *user); err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}

	return s.repo.Transaction.Delete(ctx, transaction.ID)
}

func (s *TransactionService) checkBalance(user model.User, amount float32) bool {
	if user.Balance < amount {
		return false
	}

	return true
}

func (s *TransactionService) updateBalance(user *model.User, amount float32) {
	user.Balance += amount
}
