package app

import (
	"fmt"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/config"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/service"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/storage/user/memory"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/storage/user/postgre"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport/handler"
	client "github.com/jumagaliev1/one_sdu/lecture2/hw/pkg/database/postgre"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *log.Logger) {
	cfg, err := config.New("configs/", log)
	if err != nil {
		log.Error(err)
		return
	}

	postgreClient, err := client.NewClient(cfg.Postgres)

	if err != nil {
		log.Error(err)
		return
	}

	log.Info("Successfully connect to Database")
	postgre.New(postgreClient, log)
	r := memory.New(log)
	s := service.New(r, log)
	h := handler.New(s, log)

	e := transport.Init(h)

	go GracefullyShutdown(log)
	log.Info(fmt.Sprintf("running server on %v", cfg.Server.Port))
	if err := e.Start(fmt.Sprint(":", cfg.Server.Port)); err != nil {
		log.Error(err)
	}
}

func GracefullyShutdown(log *log.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s := <-quit

	log.Info(s.String())

	os.Exit(0)
}
