package handler

import (
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BorrowHandler struct {
	service *service.Service
	logger  logger.RequestLogger
}

func NewBorrowHandler(service *service.Service, logger logger.RequestLogger) *BorrowHandler {
	return &BorrowHandler{
		service: service,
		logger:  logger,
	}
}

func (h *BorrowHandler) Create(c echo.Context) error {
	var body model.Borrow

	if err := c.Bind(&body); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	borrow, err := h.service.Borrow.Create(c.Request().Context(), body)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, borrow)
}

// GetNotReturned godoc
// @Summary      Get not Returned books for user
// @Description  Get not Returned books for user
// @ID           Get not Returned books
// @Tags         borrow
// @Accept       json
// @Produce      json
// @Success	     200  {object}  []model.UserBorrow
// @Router       /getHasBookUsers [get]
func (h *BorrowHandler) GetNotReturned(c echo.Context) error {
	borrows, err := h.service.UserBorrow.GetCurrentHaveBooks(c.Request().Context())
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, borrows)
}

// GetByLastMonth godoc
// @Summary      Get Books borrowed last month
// @Description  Get Books borrowed last month
// @ID           Get Books borrowed last month
// @Tags         borrow
// @Accept       json
// @Produce      json
// @Success	     200  {object}  []model.UserBorrow
// @Router       /getLastMonthly [get]
func (h *BorrowHandler) GetByLastMonth(c echo.Context) error {
	borrows, err := h.service.UserBorrow.GetUserBookLastMonthly(c.Request().Context())
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, borrows)
}
