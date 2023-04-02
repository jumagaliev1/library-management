package handler

import (
	"context"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
	"github.com/labstack/gommon/log"
)

type IService interface {
	Create(ctx context.Context, input model.UserInput) (*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}

type Handler struct {
	service IService
	logger  *log.Logger
}

func New(s IService, logger *log.Logger) *Handler {
	return &Handler{
		service: s,
		logger:  logger,
	}
}
