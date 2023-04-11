package main

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/one_edu/internal/app"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/logger"

	_ "github.com/jumagaliev1/one_edu/docs"
)

// @title OneLab Homework API
// @version 1.0
// @description This is a sample server for homework demo server.

// @contact.name Alibi Zhumagaliyev
// @contact.url @AZhumagaliyev
// @contact.email alibi.zhumagaliyev@gmail.com

// @host localhost:8000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description OAuth protects our entity endpoints
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New("configs/")
	if err != nil {

	}
	_ = logger.Logger(ctx)
	a := app.New(cfg, logger.Logger(ctx))
	if err := a.Run(ctx); err != nil {
		fmt.Println("FATAL")
	}
}
