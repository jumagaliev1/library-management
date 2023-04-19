package service

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/model"
)

type IUserService interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context) ([]*model.User, error)
	CheckPassword(encPass, providedPassword string) error
	HashPassword(password string) (string, error)
	Auth(ctx context.Context, user model.AuthUser) error
	RefreshToken() (string, error)
	GenerateToken(user model.AuthUser) (string, error)
	ParseToken(accessToken string) (string, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserFromRequest(ctx context.Context) (*model.User, error)
	ChangePassword(ctx context.Context, body model.PasswordReq) error
}

type IBookService interface {
	Create(ctx context.Context, book model.Book) (*model.Book, error)
	GetByTitle(ctx context.Context, title string) (*model.Book, error)
	GetAll(ctx context.Context) ([]model.Book, error)
}

type IBorrowService interface {
	Create(ctx context.Context, borrow model.Borrow) (*model.Borrow, error)
	GetAll(ctx context.Context) ([]model.Borrow, error)
	GetNoReturned(ctx context.Context) ([]model.Borrow, error)
}

type IUserBorrowService interface {
	GetCurrentBooks(ctx context.Context) ([]model.CurrentBooks, error)
	GetCurrentHaveBooks(ctx context.Context) ([]model.UserBorrow, error)
	GetUserBookLastMonthly(ctx context.Context) ([]model.UserBorrow, error)
}

type ITransactionService interface {
	Create(ctx context.Context, transaction model.Transaction) (*model.Transaction, error)
	Cancel(ctx context.Context, transactionID uint) error
}
