package app

import (
	"fmt"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/config"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/service"
	_ "github.com/jumagaliev1/one_sdu/lecture2/hw/internal/storage/user/memory"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/storage/user/postgre"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport/handler"
	client "github.com/jumagaliev1/one_sdu/lecture2/hw/pkg/database/postgre"
	"github.com/labstack/gommon/log"
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
	r2 := postgre.New(postgreClient, log)
	//r := memory.New(log)
	s := service.New(r2, log)
	h := handler.New(s, log)

	e := transport.Init(h)

	log.Info(fmt.Sprintf("running server on %v", cfg.Server.Port))
	if err := e.Start(":8080"); err != nil {
		log.Error(err)
	}

	//srv := transport.NewServer(80)

	//if err := srv.ListenAndServe(); err != nil {
	//	fmt.Print("Error in listen")
	//}
}
