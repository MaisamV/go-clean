package query

import (
	"github.com/go-clean/internal/swagger/ports"
)

// SwaggerQueryHandler handles swagger-related query operations
type SwaggerQueryHandler struct {
	swaggerProvider ports.SwaggerProvider
}

// NewSwaggerQueryHandler creates a new instance of SwaggerQueryHandler
func NewSwaggerQueryHandler(swaggerProvider ports.SwaggerProvider) *SwaggerQueryHandler {
	return &SwaggerQueryHandler{
		swaggerProvider: swaggerProvider,
	}
}

// GetOpenAPISpec returns the OpenAPI specification as JSON
func (h *SwaggerQueryHandler) GetOpenAPISpec() ([]byte, error) {
	return h.swaggerProvider.GetOpenAPISpec()
}

// GetSwaggerHTML returns the Swagger UI HTML page
func (h *SwaggerQueryHandler) GetSwaggerHTML() ([]byte, error) {
	return h.swaggerProvider.GetSwaggerHTML()
}