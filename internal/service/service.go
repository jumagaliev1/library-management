package service

import (
	"errors"
	"github.com/jumagaliev1/one_edu/internal/config"
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

func New(repo *storage.Storage, cfg config.Config) (*Service, error) {
	if repo == nil {
		return nil, errors.New("No storage")
	}
	usrService := NewUserService(repo, cfg)
	bkService := NewBookService(repo, cfg)
	borrowService := NewBorrowService(repo, cfg)
	userBorrowService := NewUserBorrowService(repo)
	transService := NewTransactionService(repo, usrService)
	return &Service{
		User:        usrService,
		Book:        bkService,
		Borrow:      borrowService,
		UserBorrow:  userBorrowService,
		Transaction: transService,
	}, nil
}
