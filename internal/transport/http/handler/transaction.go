package handler

import (
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	service *service.Service
	logger  logger.RequestLogger
}

func NewTransactionHandler(service *service.Service, logger logger.RequestLogger) *TransactionHandler {
	return &TransactionHandler{
		service: service,
		logger:  logger,
	}
}

// CreateTransaction godoc
// @Summary      Create Transaction
// @Description  Create Transaction
// @ID           CreateTransaction
// @Security	ApiKeyAuth
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        rq   body      model.TransactionReq  true  "Входящие данные"
// @Success	     200  {object}  model.Transaction
// @Router       /transaction [post]
func (h *TransactionHandler) Create(c echo.Context) error {
	var body model.TransactionReq

	if err := c.Bind(&body); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	transaction := &model.Transaction{
		Amount: body.Amount,
		BookID: body.BookID,
	}
	transaction, err := h.service.Transaction.Create(c.Request().Context(), *transaction)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transaction)
}

// CancelTransaction godoc
// @Summary      Cancel Transaction
// @Description  Cancel Transaction
// @ID           CamcelTransaction
// @Security	ApiKeyAuth
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Входящие данные"
// @Success	     200  {object}  string
// @Router       /transaction [delete]
func (h *TransactionHandler) Cancel(c echo.Context) error {
	var ID uint
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}
	ID = uint(id)

	err = h.service.Transaction.Cancel(c.Request().Context(), ID)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "cancelled")
}
