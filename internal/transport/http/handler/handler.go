package handler

import (
	"errors"
	"github.com/jumagaliev1/one_edu/internal/service"
	jwt "github.com/jumagaliev1/one_edu/internal/transport/middleware"
)

type Handler struct {
	User   IUserHandler
	Book   IBookHandler
	Borrow IBorrowHandler
}

func New(service *service.Service, jwt *jwt.JWTAuth) (*Handler, error) {
	if service == nil {
		return nil, errors.New("No given service")
	}
	usr := NewUserHandler(service, jwt)
	book := NewBookHandler(service)
	borrow := NewBorrowHandler(service)
	var handler Handler

	handler.User = usr
	handler.Book = book
	handler.Borrow = borrow

	return &handler, nil

}
