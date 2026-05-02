package api

import (
	"log"
	"net/http"

	"github.com/Kevinrestrepoh/event-simulator/event"
	"github.com/Kevinrestrepoh/event-simulator/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	event *event.Service
	log   *logger.Logger
}

func NewHandler(event *event.Service, log *logger.Logger) *Handler {
	return &Handler{event: event, log: log}
}

func (h *Handler) Run(addr string) {
	router := http.NewServeMux()

	router.HandleFunc("/events/error", h.EmitError)
	router.HandleFunc("/events/latency", h.EmitLatency)
	router.HandleFunc("/events/custom", h.EmitCustom)
	router.HandleFunc("/health", Health)

	router.Handle("/metrics", promhttp.Handler())

	RequestID(router)
	Metrics(router)
	Logging(h.log)(router)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("Sever running on port: %s", addr)
	server.ListenAndServe()
}
