package main

import (
	"github.com/Kevinrestrepoh/event-simulator/api"
	"github.com/Kevinrestrepoh/event-simulator/event"
)

func main() {
	event := event.NewService()
	svc := api.NewHandler(event)
	svc.Run(":8080")
}
