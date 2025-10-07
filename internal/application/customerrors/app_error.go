package customerrors

import (
	"errors"
	"fmt"
	"strings"
)

type ErrorType string

const (
	VALIDATION_ERROR ErrorType = "validation"
	DEPENDENCY_ERROR ErrorType = "dependency"
)

// Reserved detail keys for consistent logging/telemetry.
const (
	DetailFile      = "file"
	DetailKind      = "kind"
	DetailVersion   = "version"
	DetailComponent = "component"
	DetailExpected  = "expected"
)

// AppError is the central application error.
// - Message: stable, user-facing summary.
// - Details: structured context (use reserved keys when applicable).
// - cause: wrapped technical cause (used by errors.Is/As via Unwrap()).
type AppError struct {
	Type    ErrorType
	Message string
	Details map[string]any
	cause   error
}

func (e *AppError) Error() string {
	if e.cause == nil {
		return fmt.Sprintf("[%s] %s", e.Type, e.Message)
	}
	return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.cause)
}

// Unwrap exposes the wrapped technical cause for errors.Is/As.
func (e *AppError) Unwrap() error { return e.cause }

// NewValidationError creates a VALIDATION_ERROR and wraps the technical cause.
// Caller MUST provide a non-nil cause; it is used for errors.Is/As lookups.
// If cause is nil, a defensive placeholder is created.
func NewValidationError(message string, cause error, details map[string]any) error {
	if cause == nil {
		cause = errors.New("missing technical cause for validation error")
	}
	return &AppError{
		Type:    VALIDATION_ERROR,
		Message: message,
		Details: cloneDetails(details),
		cause:   cause,
	}
}

// NewDependencyError builds a standardized DEPENDENCY_ERROR with a wrapped cause.
func NewDependencyError(component string) error {
	cause := fmt.Errorf("dependency %q is required", component)
	return &AppError{
		Type:    DEPENDENCY_ERROR,
		Message: "Missing required dependency",
		Details: map[string]any{
			DetailComponent: component,
		},
		cause: cause,
	}
}

// NewUnsupportedVersionError builds a VALIDATION_ERROR indicating an unsupported major version.
// expectedMajors should contain allowed majors (e.g., []int{3}).
func NewUnsupportedVersionError(file, got string, expectedMajors []int) error {
	cause := fmt.Errorf("got %s, expected one of: %s", got, formatMajors(expectedMajors))
	return NewValidationError(
		"Unsupported OpenAPI version",
		cause,
		map[string]any{
			DetailFile:     file,
			DetailVersion:  got,
			DetailExpected: expectedMajors,
		},
	)
}

// formatMajors renders majors like [3,4] as "3.x, 4.x".
func formatMajors(majors []int) string {
	if len(majors) == 0 {
		return ""
	}
	parts := make([]string, len(majors))
	for i, m := range majors {
		parts[i] = fmt.Sprintf("%d.x", m)
	}
	return strings.Join(parts, ", ")
}

// cloneDetails returns a non-nil shallow copy of details.
func cloneDetails(in map[string]any) map[string]any {
	if in == nil {
		return make(map[string]any, 1)
	}
	out := make(map[string]any, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}
