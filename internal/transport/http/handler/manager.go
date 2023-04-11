package handler

import "github.com/labstack/echo/v4"

type IUserHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Auth(c echo.Context) error
	ChangePassword(c echo.Context) error
}

type IBookHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
}

type IBorrowHandler interface {
	Create(c echo.Context) error
	GetNotReturned(c echo.Context) error
	GetByLastMonth(c echo.Context) error
}
