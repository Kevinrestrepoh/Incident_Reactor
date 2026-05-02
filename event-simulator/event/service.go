package event

import (
	"context"
	"encoding/json"
	"log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Emit(ctx context.Context, e Event) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}

	// For now: log (later → EventBridge / SQS)
	log.Println(string(data))
	return nil
}
