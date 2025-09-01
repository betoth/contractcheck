package output

// Logger is an abstraction so we can swap logging implementations
type Logger interface {
	With(kv ...any) Logger
	Named(name string) Logger

	Info(msg string, kv ...any)
	Warn(msg string, kv ...any)
	Error(msg string, kv ...any)
	Debug(msg string, kv ...any)

	Sync() error
}
