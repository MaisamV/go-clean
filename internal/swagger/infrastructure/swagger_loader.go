package infrastructure

import (
	"fmt"
	"os"
)

// SwaggerLoader implements the ports.SwaggerProvider interface
type SwaggerLoader struct {
	openApiFilePath string
	swaggerFilePath string
	openapi         []byte
	swaggerHtml     []byte
}

// NewSwaggerLoader creates a new instance of DocsAdapter
func NewSwaggerLoader(openApiFilePath string, swaggerFilePath string) *SwaggerLoader {
	return &SwaggerLoader{
		openApiFilePath: openApiFilePath,
		swaggerFilePath: swaggerFilePath,
	}
}

func (a *SwaggerLoader) Init() error {
	data, err := os.ReadFile(a.openApiFilePath)
	if err != nil {
		return fmt.Errorf("failed to read OpenAPI spec: %w", err)
	}
	a.openapi = data

	html, err := os.ReadFile(a.swaggerFilePath)
	if err != nil {
		return fmt.Errorf("failed to read OpenAPI spec: %w", err)
	}
	a.swaggerHtml = html
	return nil
}

// GetOpenAPISpec loads the OpenAPI specification from file
func (a *SwaggerLoader) GetOpenAPISpec() ([]byte, error) {
	return a.openapi, nil
}

// GetSwaggerHTML generates the Swagger UI HTML
func (a *SwaggerLoader) GetSwaggerHTML() ([]byte, error) {
	return a.swaggerHtml, nil
}
