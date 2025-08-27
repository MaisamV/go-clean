package http

import (
	"github.com/go-clean/internal/swagger/application/query"
	"github.com/gofiber/fiber/v2"
)

// DocsHandler handles HTTP requests for documentation endpoints
type DocsHandler struct {
	swaggerQueryHandler *query.SwaggerQueryHandler
}

// NewDocsHandler creates a new instance of DocsHandler
func NewDocsHandler(swaggerQueryHandler *query.SwaggerQueryHandler) *DocsHandler {
	return &DocsHandler{
		swaggerQueryHandler: swaggerQueryHandler,
	}
}

// GetOpenAPISpec handles GET /api/docs/openapi.yaml
// @Summary Get OpenAPI Specification
// @Description Returns the OpenAPI specification in YAML format
// @Tags Documentation
// @Produce text/plain
// @Success 200 {string} string "OpenAPI specification in YAML format"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/docs/openapi.yaml [get]
func (h *DocsHandler) GetOpenAPISpec(c *fiber.Ctx) error {
	spec, err := h.swaggerQueryHandler.GetOpenAPISpec()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load OpenAPI specification",
		})
	}

	c.Set("Content-Type", "text/yaml")
	return c.Send(spec)
}

// GetSwaggerUI handles GET /api/docs
// @Summary Get Swagger UI
// @Description Returns the Swagger UI HTML page for API documentation
// @Tags Documentation
// @Produce text/html
// @Success 200 {string} string "Swagger UI HTML page"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/docs [get]
func (h *DocsHandler) GetSwaggerUI(c *fiber.Ctx) error {
	html, err := h.swaggerQueryHandler.GetSwaggerHTML()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate Swagger UI",
		})
	}

	c.Set("Content-Type", "text/html")
	return c.Send(html)
}

// RegisterRoutes registers the documentation routes
func (h *DocsHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/swagger", h.GetSwaggerUI)
	app.Get("/openapi.yaml", h.GetOpenAPISpec)
}
