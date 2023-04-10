package handler

import (
	"errors"
	"fmt"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	jwt "github.com/jumagaliev1/one_edu/internal/transport/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	service *service.Service
	jwt     *jwt.JWTAuth
}

func NewUserHandler(service *service.Service, jwt *jwt.JWTAuth) *UserHandler {
	return &UserHandler{
		service: service,
		jwt:     jwt,
	}
}

func (h *UserHandler) Create(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	usr, err := h.service.User.Create(c.Request().Context(), user)

	if err != nil {
		switch {
		case err == model.ErrDuplicateKey:
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *UserHandler) Auth(c echo.Context) error {
	var input model.AuthUser
	if err := c.Bind(&input); err != nil {
		return err
	}
	fmt.Println(input.Password, input.Username)
	if err := h.service.User.Auth(c.Request().Context(), input); err != nil {
		return err
	}
	token, err := h.jwt.GenerateJWT(input.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func (h *UserHandler) Get(c echo.Context) error {
	claim := c.Request().Context().Value(model.ContextUsername)
	username, ok := claim.(string)
	if !ok {
		return errors.New("cannot validate username")
	}
	user, err := h.service.User.GetByUsername(c.Request().Context(), username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
