package api

import (
	"log"
	"net/http"

	"github.com/Kevinrestrepoh/event-simulator/event"
)

type Handler struct {
	event *event.Service
}

func NewHandler(event *event.Service) *Handler {
	return &Handler{event: event}
}

func (h *Handler) Run(addr string) {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("Sever running on port: %s", addr)
	server.ListenAndServe()
}
