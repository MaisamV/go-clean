package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-clean/internal/docs/ports"
)

// docsAdapter implements the ports.DocsPort interface
type docsAdapter struct {
	apiDir string
}

// NewDocsAdapter creates a new instance of DocsAdapter
func NewDocsAdapter(apiDir string) ports.DocsPort {
	return &docsAdapter{
		apiDir: apiDir,
	}
}

// LoadOpenAPISpec loads the OpenAPI specification from file
func (a *docsAdapter) LoadOpenAPISpec() ([]byte, error) {
	specPath := filepath.Join(a.apiDir, "openapi.yaml")
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read OpenAPI spec: %w", err)
	}
	return data, nil
}

// GenerateSwaggerHTML generates the Swagger UI HTML
func (a *docsAdapter) GenerateSwaggerHTML() ([]byte, error) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui.css" />
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }
        *, *:before, *:after {
            box-sizing: inherit;
        }
        body {
            margin:0;
            background: #fafafa;
        }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '/api/docs/openapi.yaml',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
        };
    </script>
</body>
</html>`
	return []byte(html), nil
}