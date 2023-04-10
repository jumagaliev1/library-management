package handler

import (
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookHandler struct {
	service *service.Service
}

func NewBookHandler(service *service.Service) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) Create(c echo.Context) error {
	var input model.Book
	book, err := h.service.Book.Create(c.Request().Context(), input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Get(c echo.Context) error {
	var title string
	book, err := h.service.Book.GetByTitle(c.Request().Context(), title)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}
