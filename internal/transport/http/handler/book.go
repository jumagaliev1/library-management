package handler

import (
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookHandler struct {
	service *service.Service
	logger  logger.RequestLogger
}

func NewBookHandler(service *service.Service, logger logger.RequestLogger) *BookHandler {
	return &BookHandler{
		service: service,
		logger:  logger,
	}
}

// Create Book godoc
// @Summary      Create Book
// @Description  Create Book
// @ID           CreateBook
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        rq   body      model.BookReq true  "Входящие данные"
// @Success	     200  {object}  string
// @Router       /book [post]
func (h *BookHandler) Create(c echo.Context) error {
	var input model.BookReq

	if err := c.Bind(&input); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	book, err := h.service.Book.Create(c.Request().Context(), *input.MapperToBook())
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Get(c echo.Context) error {
	var title string
	book, err := h.service.Book.GetByTitle(c.Request().Context(), title)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, book)
}

// Get All Book godoc
// @Summary      Get all Book
// @Description  Get all Book
// @ID           GetAllBook
// @Tags         book
// @Accept       json
// @Produce      json
// @Success	     200  {object}  string
// @Router       /book [get]
func (h *BookHandler) GetAll(c echo.Context) error {
	books, err := h.service.Book.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, books)
}
