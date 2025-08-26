package domain

// DocsService defines the interface for documentation services
type DocsService interface {
	// GetOpenAPISpec returns the OpenAPI specification as JSON
	GetOpenAPISpec() ([]byte, error)
	
	// GetSwaggerHTML returns the Swagger UI HTML page
	GetSwaggerHTML() ([]byte, error)
}