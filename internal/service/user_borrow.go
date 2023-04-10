package service

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage"
)

type UserBorrowService struct {
	repo *storage.Storage
}

func NewUserBorrowService(repo *storage.Storage) *UserBorrowService {
	return &UserBorrowService{repo: repo}
}

func (s *UserBorrowService) GetCurrentHaveBooks(ctx context.Context) ([]model.UserBorrow, error) {
	borrows, err := s.repo.Borrow.GetNoReturned(ctx)
	if err != nil {
		return nil, err
	}

	return s.GroupBy(ctx, borrows)
}

func (s *UserBorrowService) GetUserBookLastMonthly(ctx context.Context) ([]model.UserBorrow, error) {
	borrows, err := s.repo.Borrow.GetByTime(ctx)
	if err != nil {
		return nil, err
	}

	return s.GroupBy(ctx, borrows)
}

func (s *UserBorrowService) GroupBy(ctx context.Context, borrows []model.Borrow) ([]model.UserBorrow, error) {
	currentBorrows := make(map[uint][]uint)
	for _, borrow := range borrows {
		if _, ok := currentBorrows[borrow.UserID]; !ok {
			currentBorrows[borrow.UserID] = make([]uint, 0)
		}
		currentBorrows[borrow.UserID] = append(currentBorrows[borrow.UserID], borrow.BookID)
	}

	res := make(map[model.User][]model.Book)
	for userID := range currentBorrows {
		user, err := s.repo.User.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}
		for _, bookID := range currentBorrows[userID] {
			book, err := s.repo.Book.GetByID(ctx, bookID)
			if err != nil {
				return nil, err
			}
			res[*user] = append(res[*user], *book)
		}
	}
	var userBorrow []model.UserBorrow
	for user := range res {
		userBorrow = append(userBorrow, model.UserBorrow{User: user, Books: res[user]})
	}

	return userBorrow, nil
}