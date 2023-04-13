package service

import (
	"errors"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/storage"
	"time"
)

const (
	TIME_MONTH = 30 * 24 * time.Hour
)

type Service struct {
	User        IUserService
	Book        IBookService
	Borrow      IBorrowService
	UserBorrow  IUserBorrowService
	Transaction ITransactionService
}

func New(repo *storage.Storage, cfg config.Config, logger logger.RequestLogger) (*Service, error) {
	if repo == nil {
		return nil, errors.New("No storage")
	}
	usrService := NewUserService(repo, cfg, logger)
	bkService := NewBookService(repo, cfg, logger)
	borrowService := NewBorrowService(repo, cfg, logger)
	userBorrowService := NewUserBorrowService(repo, logger)
	transService := NewTransactionService(repo, usrService, logger)
	return &Service{
		User:        usrService,
		Book:        bkService,
		Borrow:      borrowService,
		UserBorrow:  userBorrowService,
		Transaction: transService,
	}, nil
}
