package input

// VersionPolicy defines the rules used by the application to accept or reject
// an incoming spec based on its *major* version. This abstraction lets you
// decouple business policy (e.g., “only OpenAPI 3.x”) from the loader/parser.
type VersionPolicy interface {
	IsSupported(major int) bool
	SupportedVersions() []int
	FormatVersions() string
}
