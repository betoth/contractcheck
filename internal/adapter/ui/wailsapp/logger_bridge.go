// internal/adapter/ui/wailsapp/logger_bridge.go
package wailsapp

import "github.com/betoth/contractcheck/internal/application/ports/output"

// LoggerBridge exposes logging methods to the frontend (via Wails).
// It proxies calls from JS -> Go -> zap logger.
type LoggerBridge struct {
	logger output.Logger
}

// NewLoggerBridge constructs a bridge with the injected logger.
func NewLoggerBridge(logger output.Logger) *LoggerBridge {
	return &LoggerBridge{logger: logger}
}

// Info logs at info level
func (l *LoggerBridge) Info(msg string, meta string) {
	l.logger.Info(msg, "meta", meta)
}

// Warn logs at warn level
func (l *LoggerBridge) Warn(msg string, meta string) {
	l.logger.Warn(msg, "meta", meta)
}

// Error logs at error level
func (l *LoggerBridge) Error(msg string, meta string) {
	l.logger.Error(msg, "meta", meta)
}

// Debug logs at debug level
func (l *LoggerBridge) Debug(msg string, meta string) {
	l.logger.Debug(msg, "meta", meta)
}
