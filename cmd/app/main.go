package main

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/one_edu/internal/app"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/logger"
)

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
