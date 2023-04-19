package app

import (
	"context"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/service"
	"github.com/jumagaliev1/one_edu/internal/storage"
	http "github.com/jumagaliev1/one_edu/internal/transport"
	"github.com/jumagaliev1/one_edu/internal/transport/http/handler"
	"github.com/jumagaliev1/one_edu/internal/transport/middleware"
	pb "github.com/jumagaliev1/one_edu/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg    *config.Config
	logger logger.RequestLogger
}

func New(cfg *config.Config, logger logger.RequestLogger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Run(ctx context.Context) error {
	stg, err := storage.New(ctx, a.cfg, a.logger)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(a.cfg.Server.GRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	grpcServer := pb.NewTransactionServiceClient(conn)

	svc, svcErr := service.New(stg, *a.cfg, a.logger, grpcServer)
	if svcErr != nil {
		return svcErr
	}
	jwtAuth := middleware.NewJWTAuth(a.cfg, svc.User)
	h, err := handler.New(svc, jwtAuth, a.logger, grpcServer)
	if err != nil {
		return err
	}
	mdlware := middleware.Middleware{}

	HTTPServer := http.NewServer(a.cfg, h, jwtAuth, mdlware, grpcServer)

	return HTTPServer.StartHTTPServer(ctx)
}
