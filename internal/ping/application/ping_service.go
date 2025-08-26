package application

import (
	"context"

	"github.com/go-clean/internal/ping/domain"
)

// Static ping response to avoid creating new instances on every request
var staticPingResponse = &domain.PingResponse{
	Message: "PONG",
}

// PingService handles ping-related use cases
type PingService struct {
	// No dependencies needed for simple ping functionality
}

// NewPingService creates a new ping service
func NewPingService() *PingService {
	return &PingService{}
}

// Ping handles the ping use case and returns a ping response
func (s *PingService) Ping(ctx context.Context) (*domain.PingResponse, error) {
	// Return the static response for better performance
	return staticPingResponse, nil
}