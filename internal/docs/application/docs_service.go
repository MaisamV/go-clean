package application

import (
	"github.com/go-clean/internal/docs/domain"
	"github.com/go-clean/internal/docs/ports"
)

// docsService implements the domain.DocsService interface
type docsService struct {
	docsPort ports.DocsPort
}

// NewDocsService creates a new instance of DocsService
func NewDocsService(docsPort ports.DocsPort) domain.DocsService {
	return &docsService{
		docsPort: docsPort,
	}
}

// GetOpenAPISpec returns the OpenAPI specification as JSON
func (s *docsService) GetOpenAPISpec() ([]byte, error) {
	return s.docsPort.LoadOpenAPISpec()
}

// GetSwaggerHTML returns the Swagger UI HTML page
func (s *docsService) GetSwaggerHTML() ([]byte, error) {
	return s.docsPort.GenerateSwaggerHTML()
}