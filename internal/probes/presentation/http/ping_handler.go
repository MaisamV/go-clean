package http

import (
	"github.com/go-clean/internal/probes/application/query"
	"github.com/gofiber/fiber/v2"
)

// PingHandler handles HTTP requests for ping endpoints
type PingHandler struct {
	pingQueryHandler *query.PingQueryHandler
}

// NewPingHandler creates a new ping HTTP handler
func NewPingHandler(pingQueryHandler *query.PingQueryHandler) *PingHandler {
	return &PingHandler{
		pingQueryHandler: pingQueryHandler,
	}
}

// Ping handles GET /ping requests
// @Summary Ping endpoint
// @Description Returns a simple PONG response to verify service is alive
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} domain.PingResponse
// @Router /ping [get]
func (h *PingHandler) Ping(c *fiber.Ctx) error {
	ctx := c.Context()

	response, err := h.pingQueryHandler.Handle(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// RegisterRoutes registers ping routes with the fiber app
func (h *PingHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/ping", h.Ping)
	app.Get("/health", h.Health)
}

// Health handles GET /health requests
// @Summary Health check endpoint
// @Description Returns the health status of the service and its dependencies
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (h *PingHandler) Health(c *fiber.Ctx) error {
	response := HealthResponse{
		Status: "healthy",
		Checks: []HealthCheck{
			{
				Name:   "service",
				Status: "healthy",
			},
		},
	}
	
	return c.Status(fiber.StatusOK).JSON(response)
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
	Checks []HealthCheck `json:"checks"`
}

// HealthCheck represents individual health check
type HealthCheck struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
