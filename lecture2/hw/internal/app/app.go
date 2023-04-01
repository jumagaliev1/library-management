package app

import (
	"fmt"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/service"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/storage/user/memory"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/transport/handler"
)

func Run() {
	//e := echo.New()
	//fmt.Println("Run...")
	//
	//if err := e.Start(":8080"); err != nil {
	//	// TO-DO logger
	//}
	r := memory.New()
	s := service.New(r)
	h := handler.New(s)

	e := transport.Init(h)

	if err := e.Start(":8080"); err != nil {
		fmt.Println("Error server", err)
	}
	//srv := transport.NewServer(80)

	//if err := srv.ListenAndServe(); err != nil {
	//	fmt.Print("Error in listen")
	//}
}

//GetByID(ctx context.Context, id int) (*model.User, error)
//Create(ctx context.Context, input model.UserInput) (*model.User, error)
