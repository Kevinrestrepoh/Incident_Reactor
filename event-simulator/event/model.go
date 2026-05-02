package event

import "time"

type EventType string

const (
	EventError   EventType = "ERROR"
	EventLatency EventType = "LATENCY"
	EventCustom  EventType = "CUSTOM"
)

type Event struct {
	ID        string            `json:"id"`
	Type      EventType         `json:"type"`
	Service   string            `json:"service"`
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata"`
}
