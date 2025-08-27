package http

import (
	"github.com/go-clean/internal/probes/application/query"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// HealthHandler handles HTTP requests for health endpoints
type HealthHandler struct {
	healthService *query.HealthService
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(healthService *query.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

// GetHealth handles GET /health requests
// @Summary Get system health status
// @Description Returns the health status of the system including database and Redis connectivity
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} domain.HealthResponse "System is healthy"
// @Success 503 {object} domain.HealthResponse "System is unhealthy"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /health [get]
func (h *HealthHandler) GetHealth(c *fiber.Ctx) error {
	ctx := c.Context()

	// Get health status from service
	healthResponse, err := h.healthService.GetHealthStatus(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to check system health",
			"details": err.Error(),
		})
	}

	// Return appropriate HTTP status based on health
	statusCode := http.StatusOK
	if !healthResponse.IsHealthy() {
		statusCode = http.StatusServiceUnavailable
	}

	return c.Status(statusCode).JSON(healthResponse)
}

// RegisterRoutes registers health-related routes
func (h *HealthHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/health", h.GetHealth)
}
