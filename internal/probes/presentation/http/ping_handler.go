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
}
