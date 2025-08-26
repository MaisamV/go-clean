package ports

import (
	"context"

	"github.com/go-clean/internal/ping/domain"
)

// PingServicePort defines the interface for ping-related operations
// This port can be used by other modules or external systems to interact with ping functionality
type PingServicePort interface {
	// Ping executes a ping operation and returns the response
	Ping(ctx context.Context) (*domain.PingResponse, error)
}