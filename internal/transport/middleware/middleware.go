package middleware

import (
	"github.com/google/uuid"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
}

func (m *Middleware) RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		xRequestID := uuid.NewString()
		c.SetRequest(c.Request().WithContext(logger.WithRqId(c.Request().Context(), xRequestID)))
		return next(c)
	}
}
