package app

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/jumagaliev1/one_edu/internal/storage"
	http "github.com/jumagaliev1/one_edu/internal/transport"
	"github.com/jumagaliev1/one_edu/internal/transport/http/handler"
	"github.com/jumagaliev1/one_edu/internal/transport/middleware"
	"go.uber.org/zap"
)

type App struct {
	cfg    *config.Config
	logger zap.SugaredLogger
}

func New(cfg *config.Config, logger zap.SugaredLogger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Run(ctx context.Context) error {
	stg, err := storage.New(ctx, a.cfg)
	if err != nil {
		return err
	}

	svc, svcErr := service.New(stg, *a.cfg)
	if svcErr != nil {
		return svcErr
	}
	jwtAuth := middleware.NewJWTAuth(a.cfg, svc.User)
	h, err := handler.New(svc, jwtAuth)
	if err != nil {
		return err
	}

	HTTPServer := http.NewServer(a.cfg, h, jwtAuth)

	return HTTPServer.StartHTTPServer(ctx)
}
