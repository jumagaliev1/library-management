package main

import (
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/app"
	"github.com/labstack/gommon/log"
)

func main() {
	logger := log.New("DEV")
	app.Run(logger)
}
