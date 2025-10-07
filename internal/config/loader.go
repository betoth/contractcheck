// internal/config/loader.go
package config

import (
	"errors"
	"fmt"
)

// ErrConfigInvalid signals a structurally valid file whose semantic values
var ErrConfigInvalid = errors.New("invalid config")

// LoadAppConfig returns the effective application configuration.
// For now, we rely on in-process defaults (see config.Default).
func LoadAppConfig() (*AppConfig, error) {
	cfg := Default()

	if len(cfg.OpenAPI.SupportedMajors) == 0 {
		return nil, fmt.Errorf("%w: config field openapi.supported_majors: must not be empty", ErrConfigInvalid)
	}

	return &cfg, nil
}
