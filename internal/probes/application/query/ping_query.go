package query

import (
	"context"

	"github.com/go-clean/internal/probes/domain"
)

// Static ping response to avoid creating new instances on every request
var staticPingResponse = &domain.PingResponse{
	Message: "PONG",
}

// PingQueryHandler handles ping query operations
type PingQueryHandler struct {
	// No dependencies needed for simple ping functionality
}

// NewPingQueryHandler creates a new ping query handler
func NewPingQueryHandler() *PingQueryHandler {
	return &PingQueryHandler{}
}

// Handle processes the ping query and returns a ping response
func (h *PingQueryHandler) Handle(ctx context.Context) (*domain.PingResponse, error) {
	// Return the static response for better performance
	return staticPingResponse, nil
}