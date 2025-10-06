package openapi

import (
	"context"
	"strconv"
	"strings"
)

// OpenAPIDoc is the normalized representation returned by loaders/adapters.
// - JSON: canonical UTF-8 JSON of the OpenAPI document (source may be YAML/JSON).
// - Version: the declared semantic version from the `openapi` field (e.g., "3.0.3").
type OpenAPIDoc struct {
	JSON    []byte
	Version OpenAPIVersion
}

// OpenAPIVersion is a thin wrapper to provide safe helpers over the `openapi` field.
type OpenAPIVersion string

// String returns the raw version string (e.g., "3.0.3").
func (v OpenAPIVersion) String() string { return string(v) }

// Major extracts the major component (X in X.Y[.Z]).
// Returns 0 on any parsing issue to avoid panics and keep callers branch-friendly.
func (v OpenAPIVersion) Major() int {
	s := string(v)
	head, _, found := strings.Cut(s, ".")
	if !found || head == "" {
		return 0
	}
	n, err := strconv.Atoi(head)
	if err != nil || n < 0 {
		return 0
	}
	return n
}

// IsValid verifies the version follows "X.Y" or "X.Y.Z", where:
// - X > 0
// - Y >= 0
// - Z (optional) >= 0
// It avoids regexp for perf and zero allocs beyond the input.
// Returns false for empty/whitespace, missing parts, or trailing junk.
func (v OpenAPIVersion) IsValid() bool {
	s := strings.TrimSpace(string(v))
	if s == "" {
		return false
	}

	i, n := 0, len(s)
	scanDigits := func() (start, end int) {
		start = i
		for i < n && s[i] >= '0' && s[i] <= '9' {
			i++
		}
		return start, i
	}

	// major
	start, end := scanDigits()
	if end == start {
		return false
	}
	major, _ := strconv.Atoi(s[start:end])
	if major <= 0 {
		return false
	}

	// '.' between major.minor
	if i >= n || s[i] != '.' {
		return false
	}
	i++

	// minor
	start, end = scanDigits()
	if end == start {
		return false
	}

	// optional ".patch"
	if i == n {
		return true
	}
	if s[i] != '.' {
		return false
	}
	i++
	start, end = scanDigits()
	return end == n && end > start
}

// Loader is the output port for fetching an OpenAPI document from disk.
// Implementations should:
type Loader interface {
	Load(ctx context.Context, filePath string) (OpenAPIDoc, error)
}
