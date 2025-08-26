package application

import (
	"context"

	"github.com/go-clean/internal/ping/domain"
)

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
	// Simple ping logic - just return the standard response
	return domain.NewPingResponse(), nil
}