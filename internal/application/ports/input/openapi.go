package input

import (
	"context"

	"github.com/betoth/contractcheck/internal/application/ports/output/openapi"
)

// ImportOpenAPISpec defines the input port (use case) responsible for ingesting
// an OpenAPI specification from a file path and returning a normalized document.
type ImportOpenAPISpec interface {
	Import(ctx context.Context, filePath string) (openapi.OpenAPIDoc, error)
}
