package handler

import (
	"fmt"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	pb "github.com/jumagaliev1/one_edu/proto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	service *service.Service
	grpc    pb.TransactionServiceClient
	logger  logger.RequestLogger
}

func NewTransactionHandler(grpc pb.TransactionServiceClient, logger logger.RequestLogger, service *service.Service) *TransactionHandler {
	return &TransactionHandler{
		grpc:    grpc,
		logger:  logger,
		service: service,
	}
}

// CreateTransaction godoc
// @Summary      Create Transaction
// @Description  Create Transaction
// @ID           CreateTransaction
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

	req := &pb.CreateTransRequest{
		Transaction: &pb.Transaction{
			UserId: uint32(body.UserID),
			ItemId: uint32(body.BookID),
			Price:  int32(body.Amount),
		}}

	trans, err := h.grpc.Create(c.Request().Context(), req)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	borrow := model.Borrow{
		BookID:   body.BookID,
		UserID:   body.UserID,
		Returned: nil,
	}

	_, err = h.service.Borrow.Create(c.Request().Context(), borrow)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprint("transactionID:", trans.Id))
}

// CancelTransaction godoc
// @Summary      Cancel Transaction
// @Description  Cancel Transaction
// @ID           CamcelTransaction
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
	req := &pb.CancelTransRequest{TransactionID: int32(ID)}

	_, err = h.grpc.Cancel(c.Request().Context(), req)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "cancelled")
}

// Increment Balance User godoc
// @Summary      Increment Balance
// @Description  Increment Balance
// @ID           IncrementBalance
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        rq   body      model.IncrementBalanceReq true  "Входящие данные"
// @Success	     200  {object}  string
// @Router       /balance [post]
func (h *TransactionHandler) IncrementBalance(c echo.Context) error {
	var body model.IncrementBalanceReq

	if err := c.Bind(&body); err != nil {
		return err
	}

	req := &pb.BalanceRequest{
		UserId: int32(body.UserID),
		Amount: int32(body.Amount),
	}

	_, err := h.grpc.IncrementBalance(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}
