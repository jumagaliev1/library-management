package http

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/transport/http/handler"
	jwt "github.com/jumagaliev1/one_edu/internal/transport/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type Server struct {
	cfg     *config.Config
	App     *echo.Echo
	handler *handler.Handler
	jwt     *jwt.JWTAuth
}

func NewServer(cfg *config.Config, handler *handler.Handler, jwt *jwt.JWTAuth) *Server {
	return &Server{cfg: cfg, handler: handler, jwt: jwt}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	//c := jaegertracing.New(s.App, nil)
	//defer c.Close()
	s.SetupRoutes()

	go func() {
		if err := s.App.Start(fmt.Sprintf(":%v", s.cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:#{err}\n")
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:#{err}")
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info(map[string]interface{}{"URI": v.URI, "status": v.Status})

			return nil
		},
	}))

	//l := logger.Logger{context.Background()}

	return e
}
