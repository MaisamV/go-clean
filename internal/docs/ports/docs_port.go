package ports

// DocsPort defines the interface for documentation operations
type DocsPort interface {
	// LoadOpenAPISpec loads the OpenAPI specification from file
	LoadOpenAPISpec() ([]byte, error)
	
	// GenerateSwaggerHTML generates the Swagger UI HTML
	GenerateSwaggerHTML() ([]byte, error)
}