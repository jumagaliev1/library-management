package storage

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage/postgre"
	"gorm.io/gorm"
)

type IBookRepository interface {
	Create(ctx context.Context, book model.Book) (*model.Book, error)
	GetByTitle(ctx context.Context, title string) (*model.Book, error)
	GetByAuthor(ctx context.Context, author string) (*model.Book, error)
	GetByID(ctx context.Context, ID uint) (*model.Book, error)
}

type IUserRepository interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, ID int) error
	GetAll(ctx context.Context) ([]*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByID(ctx context.Context, ID uint) (*model.User, error)
}

type IBorrowRepository interface {
	Create(ctx context.Context, borrow model.Borrow) (*model.Borrow, error)
	GetAll(ctx context.Context) ([]model.Borrow, error)
	GetNoReturned(ctx context.Context) ([]model.Borrow, error)
	GetByTime(ctx context.Context) ([]model.Borrow, error)
}

type Storage struct {
	pg     *gorm.DB
	Book   IBookRepository
	User   IUserRepository
	Borrow IBorrowRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	uRepo := postgre.NewUserRepository(pgDB)
	bRepo := postgre.NewBookRepository(pgDB)
	borrowRepo := postgre.NewBorrowRepository(pgDB)
	var storage Storage
	storage.User = uRepo
	storage.Book = bRepo
	storage.Borrow = borrowRepo

	return &storage, nil
}
