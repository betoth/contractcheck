package openapi

import "github.com/betoth/contractcheck/internal/application/customerrors"

// ErrorKind classifies validation failures for OpenAPI import/validation flows.
type ErrorKind string

const (
	FILE_NOT_FOUND           ErrorKind = "file_not_found"
	PERMISSION_DENIED        ErrorKind = "permission_denied"
	INVALID_SYNTAX           ErrorKind = "invalid_syntax"
	EXTERNAL_REF_NOT_ALLOWED ErrorKind = "external_ref_not_allowed"
	INVALID_SPEC             ErrorKind = "invalid_spec"
	INVALID_VERSION_FORMAT   ErrorKind = "invalid_version_format"
)

// NewValidationError wraps a technical cause and returns a standardized validation error.
// - kind: machine-readable classification
// - message: stable, user-facing summary
// - file: source file path for context (optional but recommended)
// - cause: required technical cause (used by errors.Is/As)
func NewValidationError(kind ErrorKind, message, file string, cause error) error {
	return customerrors.NewValidationError(
		message,
		cause,
		map[string]any{
			customerrors.DetailKind: string(kind),
			customerrors.DetailFile: file,
		},
	)
}
