package config

// AppConfig holds user/application configuration loaded from YAML.
type AppConfig struct {
	OpenAPI OpenAPIConfig `yaml:"openapi"`
}

// OpenAPIConfig configures OpenAPI-related behavior across the app.
// SupportedMajors lists accepted major versions (e.g., 3 -> 3.x).
type OpenAPIConfig struct {
	SupportedMajors []int `yaml:"supported_majors"`
}

// Default returns a safe, opinionated configuration used on first run
// or as embedded fallback when no user config is present.
func Default() AppConfig {
	return AppConfig{
		OpenAPI: OpenAPIConfig{
			SupportedMajors: []int{3},
		},
	}
}
