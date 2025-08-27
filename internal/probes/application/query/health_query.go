package query

import (
	"context"

	"github.com/go-clean/internal/probes/domain"
	"github.com/go-clean/internal/probes/ports"
)

// GetHealthQuery represents a query to get system health status
type GetHealthQuery struct{}

// GetHealthQueryHandler handles health status queries
type GetHealthQueryHandler struct {
	databaseChecker ports.DatabaseChecker
	redisChecker    ports.RedisChecker
}

// NewGetHealthQueryHandler creates a new health query handler
func NewGetHealthQueryHandler(
	databaseChecker ports.DatabaseChecker,
	redisChecker ports.RedisChecker,
) *GetHealthQueryHandler {
	return &GetHealthQueryHandler{
		databaseChecker: databaseChecker,
		redisChecker:    redisChecker,
	}
}

// Handle executes the health check query
func (h *GetHealthQueryHandler) Handle(ctx context.Context, query GetHealthQuery) (*domain.HealthResponse, error) {
	response := domain.NewHealthResponse()

	// Check database connectivity
	if h.databaseChecker != nil {
		dbHealthy, dbResponseTime, err := h.databaseChecker.CheckDatabase(ctx)
		if err != nil {
			response.AddCheck("database", domain.CheckStatusDown, 0)
		} else {
			status := domain.CheckStatusUp
			if !dbHealthy {
				status = domain.CheckStatusDown
			}
			response.AddCheck("database", status, dbResponseTime.Milliseconds())
		}
	}

	// Check Redis connectivity
	if h.redisChecker != nil {
		redisHealthy, redisResponseTime, err := h.redisChecker.CheckRedis(ctx)
		if err != nil {
			response.AddCheck("redis", domain.CheckStatusDown, 0)
		} else {
			status := domain.CheckStatusUp
			if !redisHealthy {
				status = domain.CheckStatusDown
			}
			response.AddCheck("redis", status, redisResponseTime.Milliseconds())
		}
	}

	// Determine overall status
	response.DetermineOverallStatus()

	return response, nil
}

// HealthService implements the HealthService port
type HealthService struct {
	queryHandler *GetHealthQueryHandler
}

// NewHealthService creates a new health service
func NewHealthService(queryHandler *GetHealthQueryHandler) *HealthService {
	return &HealthService{
		queryHandler: queryHandler,
	}
}

// GetHealthStatus returns the current health status
func (s *HealthService) GetHealthStatus(ctx context.Context) (*domain.HealthResponse, error) {
	return s.queryHandler.Handle(ctx, GetHealthQuery{})
}
