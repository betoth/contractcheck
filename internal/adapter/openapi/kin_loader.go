package openapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/betoth/contractcheck/internal/application/customerrors"
	"github.com/betoth/contractcheck/internal/application/ports/output/openapi"
	"github.com/getkin/kin-openapi/openapi3"
)

// KinLoader wraps kin-openapi's Loader to conform to our output port (openapi.Loader)
// and to centralize validation + error normalization policies used by ContractCheck.
//
// Design notes:
//   - This adapter converts third-party errors into our domain AppError consistently.
//   - It validates the spec with kin-openapi before returning it upstream.
//   - It always returns the spec serialized as UTF-8 JSON ([]byte) plus the declared version,
//     leaving higher layers free to persist or further transform as needed.
type KinLoader struct {
	loader *openapi3.Loader
}

// KinLoaderOption is a functional option that allows callers to tweak the underlying
type KinLoaderOption func(*openapi3.Loader)

// NewKinLoader builds a KinLoader with safe defaults.
// By default, external $ref are NOT allowed (security/portability reasons).
// Callers may relax this via WithExternalRefsAllowed().
func NewKinLoader(opts ...KinLoaderOption) *KinLoader {
	ldr := openapi3.NewLoader()
	ldr.IsExternalRefsAllowed = false

	for _, opt := range opts {
		opt(ldr)
	}

	return &KinLoader{loader: ldr}
}

// WithExternalRefsAllowed enables resolution of external $ref.
func WithExternalRefsAllowed() KinLoaderOption {
	return func(l *openapi3.Loader) {
		l.IsExternalRefsAllowed = true
	}
}

// Load reads and validates an OpenAPI file located at filePath.
func (kin *KinLoader) Load(ctx context.Context, filePath string) (openapi.OpenAPIDoc, error) {
	doc, err := kin.loader.LoadFromFile(filePath)
	if err != nil {
		return openapi.OpenAPIDoc{}, kin.normalizeError(filePath, err)
	}

	if err := kin.validateDoc(ctx, doc, filePath); err != nil {
		return openapi.OpenAPIDoc{}, kin.normalizeError(filePath, err)
	}

	raw, err := json.Marshal(doc)
	if err != nil {
		return openapi.OpenAPIDoc{}, kin.normalizeError(filePath, err)
	}

	return openapi.OpenAPIDoc{
		JSON:    raw,
		Version: openapi.OpenAPIVersion(doc.OpenAPI),
	}, nil
}

// validateDoc centralizes structural validation and version checks.
func (kin *KinLoader) validateDoc(ctx context.Context, doc *openapi3.T, filePath string) error {
	if err := doc.Validate(ctx); err != nil {
		return err
	}

	if !openapi.OpenAPIVersion(doc.OpenAPI).IsValid() {
		return &customerrors.AppError{
			Type:    customerrors.VALIDATION_ERROR,
			Message: "Invalid OpenAPI version format",
			Cause:   fmt.Sprintf("got %q, expected format X.Y[.Z]", doc.OpenAPI),
			Details: map[string]interface{}{
				"file": filePath,
				"kind": openapi.INVALID_VERSION_FORMAT,
			},
		}
	}

	return nil
}

// normalizeError maps heterogeneous errors (os, YAML/JSON parsing, kin-openapi)
// into our stable AppError shape with a machine-friendly "kind".
//
// This function MUST remain deterministic: callers rely on "kind" for UX messages,
// branching logic and telemetry dashboards.
func (kin *KinLoader) normalizeError(filePath string, err error) error {
	var ae *customerrors.AppError
	if errors.As(err, &ae) {
		if ae.Details == nil {
			ae.Details = make(map[string]interface{})
		}
		ae.Details["file"] = filePath
		return ae
	}

	msg := err.Error()

	switch {
	case errors.Is(err, os.ErrNotExist):
		return &customerrors.AppError{
			Type:    customerrors.VALIDATION_ERROR,
			Message: "File not found",
			Cause:   msg,
			Details: map[string]interface{}{
				"file": filePath,
				"kind": openapi.FILE_NOT_FOUND,
			},
		}
	case errors.Is(err, os.ErrPermission):
		return &customerrors.AppError{
			Type:    customerrors.VALIDATION_ERROR,
			Message: "Permission denied",
			Cause:   msg,
			Details: map[string]interface{}{
				"file": filePath,
				"kind": openapi.PERMISSION_DENIED,
			},
		}
	case strings.Contains(msg, "yaml:") || strings.Contains(msg, "json:"):
		return &customerrors.AppError{
			Type:    customerrors.VALIDATION_ERROR,
			Message: "Invalid YAML/JSON syntax",
			Cause:   msg,
			Details: map[string]interface{}{
				"file": filePath,
				"kind": openapi.INVALID_SYNTAX,
			},
		}
	case strings.Contains(msg, "external reference"):
		return &customerrors.AppError{
			Type:    customerrors.VALIDATION_ERROR,
			Message: "External references are not allowed",
			Cause:   msg,
			Details: map[string]interface{}{
				"file": filePath,
				"kind": openapi.EXTERNAL_REF_NOT_ALLOWED,
			},
		}
	default:
		return &customerrors.AppError{
			Type:    customerrors.VALIDATION_ERROR,
			Message: "Invalid OpenAPI specification",
			Cause:   msg,
			Details: map[string]interface{}{
				"file": filePath,
				"kind": openapi.INVALID_SPEC,
			},
		}
	}
}

// Ensure KinLoader implements openapi.Loader interface
var _ openapi.Loader = (*KinLoader)(nil)
