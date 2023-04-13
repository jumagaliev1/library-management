package handler

import (
	"errors"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/service"
	jwt "github.com/jumagaliev1/one_edu/internal/transport/middleware"
)

type Handler struct {
	User        IUserHandler
	Book        IBookHandler
	Borrow      IBorrowHandler
	Transaction ITransactionHandler
}

func New(service *service.Service, jwt *jwt.JWTAuth, logger logger.RequestLogger) (*Handler, error) {
	if service == nil {
		return nil, errors.New("No given service")
	}
	usr := NewUserHandler(service, jwt, logger)
	book := NewBookHandler(service, logger)
	borrow := NewBorrowHandler(service, logger)
	trans := NewTransactionHandler(service, logger)
	var handler Handler

	handler.User = usr
	handler.Book = book
	handler.Borrow = borrow
	handler.Transaction = trans

	return &handler, nil

}
