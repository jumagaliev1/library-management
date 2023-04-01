package transport

import (
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport/handler"
	"github.com/labstack/echo/v4"
)

func Init(h *handler.Handler) *echo.Echo {
	e := echo.New()

	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)

	return e
}
