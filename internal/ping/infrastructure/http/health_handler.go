package http

import (
	"github.com/gofiber/fiber/v2"
)

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

// Health handles GET /health requests
// @Summary Health check endpoint
// @Description Returns the health status of the service and its dependencies
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func Health(c *fiber.Ctx) error {
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