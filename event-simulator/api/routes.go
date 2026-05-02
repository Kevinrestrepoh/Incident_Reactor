package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Kevinrestrepoh/event-simulator/event"
	"github.com/google/uuid"
)

func (h *Handler) EmitError(w http.ResponseWriter, r *http.Request) {
	e := event.Event{
		ID:        uuid.NewString(),
		Type:      event.EventError,
		Service:   "dummy-service",
		Timestamp: time.Now(),
		Metadata: map[string]string{
			"message": "simulated error",
		},
	}

	if err := h.event.Emit(r.Context(), e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) EmitLatency(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)

	e := event.Event{
		ID:        uuid.NewString(),
		Type:      event.EventLatency,
		Service:   "dummy-service",
		Timestamp: time.Now(),
		Metadata: map[string]string{
			"duration": "2s",
		},
	}

	if err := h.event.Emit(r.Context(), e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) EmitCustom(w http.ResponseWriter, r *http.Request) {
	var meta map[string]string
	if err := json.NewDecoder(r.Body).Decode(&meta); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e := event.Event{
		ID:        uuid.NewString(),
		Type:      event.EventCustom,
		Service:   "custom",
		Timestamp: time.Now(),
		Metadata:  meta,
	}

	if err := h.event.Emit(r.Context(), e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
