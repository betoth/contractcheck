package logger

import (
	"github.com/betoth/contractcheck/internal/application/ports/output"
	"go.uber.org/zap"
)

// ZapLogger adapts zap.SugaredLogger to the output.Logger port.
type ZapLogger struct {
	*zap.SugaredLogger
}

// New returns an output.Logger. Falling back to zap.NewNop() on error keeps boot safe.
func New() output.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		return &ZapLogger{zap.NewNop().Sugar()}
	}
	return &ZapLogger{l.Sugar()}
}

// With returns a derived logger with extra structured fields.
func (l *ZapLogger) With(kv ...any) output.Logger {
	return &ZapLogger{l.SugaredLogger.With(kv...)}
}

// Named returns a derived logger with a sub-scope name.
func (l *ZapLogger) Named(name string) output.Logger {
	return &ZapLogger{l.SugaredLogger.Named(name)}
}

func (l *ZapLogger) Info(msg string, kv ...any)  { l.Infow(msg, kv...) }
func (l *ZapLogger) Warn(msg string, kv ...any)  { l.Warnw(msg, kv...) }
func (l *ZapLogger) Error(msg string, kv ...any) { l.Errorw(msg, kv...) }
func (l *ZapLogger) Debug(msg string, kv ...any) { l.Debugw(msg, kv...) }
func (l *ZapLogger) Sync() error                 { return l.SugaredLogger.Sync() }
