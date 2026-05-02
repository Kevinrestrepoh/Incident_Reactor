package main

import (
	"github.com/Kevinrestrepoh/event-simulator/api"
	"github.com/Kevinrestrepoh/event-simulator/event"
	"github.com/Kevinrestrepoh/event-simulator/logger"
)

func main() {
	event := event.NewService()
	logger := logger.New()

	svc := api.NewHandler(event, logger)
	svc.Run(":8080")
}
