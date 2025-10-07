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
		// Let normalizeError map this; keep raw error as cause.
		return err
	}

	if !openapi.OpenAPIVersion(doc.OpenAPI).IsValid() {
		cause := fmt.Errorf("got %q, expected format X.Y[.Z]", doc.OpenAPI)

		// Prefer the port helper to keep taxonomy centralized.
		err := openapi.NewValidationError(
			openapi.INVALID_VERSION_FORMAT,
			"Invalid OpenAPI version format",
			filePath,
			cause,
		)

		// Enrich with the declared version for easier troubleshooting.
		var ae *customerrors.AppError
		if errors.As(err, &ae) {
			if ae.Details == nil {
				ae.Details = make(map[string]any)
			}
			ae.Details[customerrors.DetailVersion] = doc.OpenAPI
		}
		return err
	}

	return nil
}

// normalizeError maps heterogeneous errors (context, os, YAML/JSON parsing, kin-openapi)
// into our stable AppError shape with a machine-friendly "kind".
//
// This function MUST remain deterministic: callers rely on "kind" for UX messages,
// branching logic and telemetry dashboards.
func (kin *KinLoader) normalizeError(filePath string, err error) error {
	// Pass through context cancellation/timeouts unchanged — they are control-flow signals.
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	// If it's already our AppError, just enrich file (if missing) and bubble up.
	var ae *customerrors.AppError
	if errors.As(err, &ae) {
		if ae.Details == nil {
			ae.Details = make(map[string]any)
		}
		if _, ok := ae.Details[customerrors.DetailFile]; !ok {
			ae.Details[customerrors.DetailFile] = filePath
		}
		return ae
	}

	switch {
	case errors.Is(err, os.ErrNotExist):
		return customerrors.NewValidationError(
			"File not found",
			err,
			map[string]any{
				customerrors.DetailFile: filePath,
				customerrors.DetailKind: openapi.FILE_NOT_FOUND,
			},
		)

	case errors.Is(err, os.ErrPermission):
		return customerrors.NewValidationError(
			"Permission denied",
			err,
			map[string]any{
				customerrors.DetailFile: filePath,
				customerrors.DetailKind: openapi.PERMISSION_DENIED,
			},
		)

	default:
		// Heuristics for YAML/JSON parse errors (based on typical vendor messages).
		msg := err.Error()
		if strings.Contains(msg, "yaml:") || strings.Contains(msg, "json:") {
			return customerrors.NewValidationError(
				"Invalid YAML/JSON syntax",
				err,
				map[string]any{
					customerrors.DetailFile: filePath,
					customerrors.DetailKind: openapi.INVALID_SYNTAX,
				},
			)
		}
		if strings.Contains(msg, "external reference") {
			return customerrors.NewValidationError(
				"External references are not allowed",
				err,
				map[string]any{
					customerrors.DetailFile: filePath,
					customerrors.DetailKind: openapi.EXTERNAL_REF_NOT_ALLOWED,
				},
			)
		}

		// Catch-all for any other validation/parse/semantic problem.
		return customerrors.NewValidationError(
			"Invalid OpenAPI specification",
			err,
			map[string]any{
				customerrors.DetailFile: filePath,
				customerrors.DetailKind: openapi.INVALID_SPEC,
			},
		)
	}
}

// Ensure KinLoader implements openapi.Loader interface
var _ openapi.Loader = (*KinLoader)(nil)
