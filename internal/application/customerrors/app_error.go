package customerrors

import (
	"fmt"
	"strings"
)

// ErrorType classifies an application error so callers can branch on intent
// (e.g., validation vs. infrastructure/dependency failures).
type ErrorType string

const (
	// VALIDATION_ERROR indicates user/input/spec validation issues that the
	// caller can typically fix (bad format, unsupported values, etc.).
	VALIDATION_ERROR ErrorType = "validation"

	// DEPENDENCY_ERROR indicates a missing or misconfigured component the app
	// depends on (adapters, services, environment), not user input.
	DEPENDENCY_ERROR ErrorType = "dependency"
)

// AppError represents a structured domain error used across the application.
type AppError struct {
	Type    ErrorType
	Message string
	Cause   string
	Details map[string]interface{}
}

// Error satisfies the error interface, producing a concise, readable summary.
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Type, e.Message, e.Cause)
}

// NewDependencyError builds a standardized DEPENDENCY_ERROR for a missing
// or required external component.
func NewDependencyError(component string) error {
	return &AppError{
		Type:    DEPENDENCY_ERROR,
		Message: "Missing required dependency",
		Cause:   fmt.Sprintf("dependency '%s' is required", component),
	}
}

// NewUnsupportedVersionError builds a VALIDATION_ERROR indicating that an
// input (e.g., OpenAPI file) declares an unsupported major version.
// expectedMajors should contain allowed major versions (e.g., []int{3}).
func NewUnsupportedVersionError(file, got string, expectedMajors []int) error {
	return &AppError{
		Type:    VALIDATION_ERROR,
		Message: "Unsupported OpenAPI version",
		Cause:   fmt.Sprintf("got %s, expected one of: %s", got, formatMajors(expectedMajors)),
		Details: map[string]interface{}{
			"file":     file,
			"version":  got,
			"expected": expectedMajors,
		},
	}
}

// formatMajors renders majors like [3,4] as "3.x, 4.x" for user-friendly messages.
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
