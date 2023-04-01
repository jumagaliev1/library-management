package handler

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) CreateUser(c echo.Context) error {
	input := new(model.UserInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(input)
	usr, err := h.service.Create(ctx, *input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(usr)
	return c.JSON(http.StatusOK, usr)

}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TO-DO logger
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	usr, err := h.service.GetByID(ctx, id)
	if err != nil {
		// TO-DO logger
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, usr)
}
