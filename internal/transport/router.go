package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) SetupRoutes() *echo.Group {
	v1 := s.App.Group("/api/v1")

	s.App.GET("/ready", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	s.App.GET("/live", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	s.App.POST("/createUser", s.handler.User.Create)
	s.App.GET("/getUser", s.handler.User.Get, s.jwt.ValidateAuth)
	s.App.POST("/auth", s.handler.User.Auth)

	s.App.GET("/getHasBookUsers", s.handler.Borrow.GetNotReturned)
	s.App.GET("/getLastMonthly", s.handler.Borrow.GetByLastMonth)
	return v1
}
