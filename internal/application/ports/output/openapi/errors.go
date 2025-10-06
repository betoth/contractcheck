package openapi

import "github.com/betoth/contractcheck/internal/application/customerrors"

// ErrorKind classifies validation failures for OpenAPI import/validation flows.
// Keep values stable: UI, telemetry, and branching logic rely on these identifiers.
type ErrorKind string

const (
	// FILE_NOT_FOUND indicates the path does not exist on disk.
	FILE_NOT_FOUND ErrorKind = "file_not_found"

	// PERMISSION_DENIED indicates the process lacks permissions to read the file.
	PERMISSION_DENIED ErrorKind = "permission_denied"

	// INVALID_SYNTAX indicates YAML/JSON parsing errors.
	INVALID_SYNTAX ErrorKind = "invalid_syntax"

	// EXTERNAL_REF_NOT_ALLOWED indicates a policy violation for external $ref usage.
	EXTERNAL_REF_NOT_ALLOWED ErrorKind = "external_ref_not_allowed"

	// INVALID_SPEC indicates a semantically invalid OpenAPI document (failed validation).
	INVALID_SPEC ErrorKind = "invalid_spec"

	// INTERNAL_ERROR is a catch-all for unexpected failures not attributable to user input.
	// Prefer more specific kinds whenever possible.
	INTERNAL_ERROR ErrorKind = "internal_error"

	// INVALID_VERSION_FORMAT indicates an invalid "openapi" version string (e.g., not X.Y[.Z]).
	INVALID_VERSION_FORMAT ErrorKind = "invalid_version_format"
)

// NewValidationError creates a structured VALIDATION_ERROR using our domain error type.
// - kind: machine-readable classification (use one of the constants above)
// - message: short human-readable summary (safe to show to users)
// - file: source file path for context (optional but recommended)
// - cause: low-level detail for diagnostics (keep technical; not required for UX)
func NewValidationError(kind ErrorKind, message string, file string, cause string) error {
	return &customerrors.AppError{
		Type:    customerrors.VALIDATION_ERROR,
		Message: message,
		Cause:   cause,
		Details: map[string]interface{}{
			"kind": kind,
			"file": file,
		},
	}
}
