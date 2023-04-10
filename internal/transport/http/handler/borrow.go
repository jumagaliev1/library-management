package handler

import (
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BorrowHandler struct {
	service *service.Service
}

func NewBorrowHandler(service *service.Service) *BorrowHandler {
	return &BorrowHandler{service: service}
}

func (h *BorrowHandler) Create(c echo.Context) error {
	var body model.Borrow

	if err := c.Bind(&body); err != nil {
		return err
	}

	borrow, err := h.service.Borrow.Create(c.Request().Context(), body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, borrow)
}

func (h *BorrowHandler) GetNotReturned(c echo.Context) error {
	borrows, err := h.service.UserBorrow.GetCurrentHaveBooks(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, borrows)
}

func (h *BorrowHandler) GetByLastMonth(c echo.Context) error {
	borrows, err := h.service.UserBorrow.GetUserBookLastMonthly(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, borrows)
}
