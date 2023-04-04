package main

import (
	"fmt"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/app"
	"github.com/labstack/gommon/log"
	"os"
	"time"
)

const ENVIRONMENT = "DEV"

func main() {
	logger := log.New(ENVIRONMENT)
	f, err := os.OpenFile(fmt.Sprint("logs/", time.Now()), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logger.Error("Error opening file")
	}

	logger.SetOutput(f)
	app.Run(logger)
	if err = f.Close(); err != nil {
		logger.Error(err)
	}
}
