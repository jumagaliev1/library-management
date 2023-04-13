package http

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func (s *Server) SetupRoutes() *echo.Group {
	v1 := s.App.Group("/api/v1", s.middleware.RequestID)
	s.App.GET("/ready", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	s.App.GET("/live", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	v1.POST("/user", s.handler.User.Create)
	v1.GET("/user", s.handler.User.Get, s.jwt.ValidateAuth)
	v1.POST("/user/password", s.handler.User.ChangePassword, s.jwt.ValidateAuth)

	v1.POST("/auth", s.handler.User.Auth)

	v1.GET("/getHasBookUsers", s.handler.Borrow.GetNotReturned)
	v1.GET("/getLastMonthly", s.handler.Borrow.GetByLastMonth)

	v1.POST("/transaction", s.handler.Transaction.Create, s.jwt.ValidateAuth)
	v1.DELETE("/transaction/:id", s.handler.Transaction.Cancel, s.jwt.ValidateAuth)

	s.App.GET("/swagger/*", echoSwagger.WrapHandler)

	return v1
}
