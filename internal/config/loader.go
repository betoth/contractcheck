// internal/config/loader.go
package config

import (
	"errors"
	"fmt"
	"sort"
)

// ErrConfigInvalid is a sentinel error indicating the configuration failed validation.
var ErrConfigInvalid = errors.New("invalid config")

// LoadAppConfig returns the effective application configuration.
func LoadAppConfig() (*AppConfig, error) {
	cfg := Default()

	if err := normalizeAndValidate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// normalizeAndValidate applies canonicalization (dedup/sort) and validates invariants.
func normalizeAndValidate(cfg *AppConfig) error {
	cfg.OpenAPI.SupportedMajors = normalizeMajors(cfg.OpenAPI.SupportedMajors)

	if len(cfg.OpenAPI.SupportedMajors) == 0 {
		return fieldErr("openapi.supported_majors", "must not be empty")
	}

	for _, m := range cfg.OpenAPI.SupportedMajors {
		if m <= 0 {
			return fieldErr("openapi.supported_majors", "must contain only positive integers (e.g., 3 for 3.x)")
		}
	}

	return nil
}

// fieldErr wraps ErrConfigInvalid with a field-specific, actionable message.
func fieldErr(field, msg string) error {
	return fmt.Errorf("%w: field %q %s", ErrConfigInvalid, field, msg)
}

// normalizeMajors removes duplicates and non-positive values, then sorts ascending.
func normalizeMajors(in []int) []int {
	if len(in) == 0 {
		return in
	}
	seen := make(map[int]struct{}, len(in))
	out := make([]int, 0, len(in))
	for _, v := range in {
		if v <= 0 {
			continue
		}
		if _, dup := seen[v]; dup {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	sort.Ints(out)
	return out
}
