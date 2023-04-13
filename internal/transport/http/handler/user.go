package handler

import (
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	jwt "github.com/jumagaliev1/one_edu/internal/transport/middleware"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
	"net/http"
)

type UserHandler struct {
	service *service.Service
	jwt     *jwt.JWTAuth
	logger  logger.RequestLogger
}

func NewUserHandler(service *service.Service, jwt *jwt.JWTAuth, logger logger.RequestLogger) *UserHandler {
	return &UserHandler{
		service: service,
		jwt:     jwt,
		logger:  logger,
	}
}

// CreateUser godoc
// @Summary      Создание пользователя
// @Description  Создание пользователя
// @ID           CreateUser
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.User  true  "Входящие данные"
// @Success	     200  {object}  model.User
// @Router       /user [post]
func (h *UserHandler) Create(c echo.Context) error {
	h.logger.Logger(c.Request().Context()).Info("creating user...")
	var user model.User
	if err := c.Bind(&user); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}
	usr, err := h.service.User.Create(c.Request().Context(), user)

	if err != nil {
		switch {
		case err == model.ErrDuplicateKey:
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
		}
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

// Auth godoc
// @Summary      Auth get JWT token
// @Description Auth get JWT token
// @ID           AuthUser
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.AuthUser  true  "Входящие данные"
// @Success	     200  {object}  string
// @Router       /auth [post]
func (h *UserHandler) Auth(c echo.Context) error {
	var input model.AuthUser
	if err := c.Bind(&input); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}
	if err := h.service.User.Auth(c.Request().Context(), input); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}
	token, err := h.jwt.GenerateJWT(input.Username)
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, token)
}

// Get User godoc
// @Summary      Get User
// @Description  Get User
// @Security	ApiKeyAuth
// @ID           GetUser
// @Tags         user
// @Accept       json
// @Produce      json
// @Success	     200  {object}  model.User
// @Router       /user [get]
func (h *UserHandler) Get(c echo.Context) error {
	user, err := h.service.User.GetUserFromRequest(c.Request().Context())
	if err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// ChangePasswordUser godoc
// @Summary      Change Password for user
// @Description  Change Passowrd for user
// @Security	ApiKeyAuth
// @ID           ChangePassword
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        rq   body      model.PasswordReq  true  "Входящие данные"
// @Success	     200  {object}  model.User
// @Router       /user/password [post]
func (h *UserHandler) ChangePassword(c echo.Context) error {
	var body model.PasswordReq

	if err := c.Bind(&body); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return err
	}

	if err := h.service.User.ChangePassword(c.Request().Context(), body); err != nil {
		h.logger.Logger(c.Request().Context()).Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "success")
}
