package logger

import (
	"github.com/betoth/contractcheck/internal/application/ports/output"
	"go.uber.org/zap"
)

type ZapLogger struct {
	*zap.SugaredLogger
}

func New() output.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.Encoding = "console"

	l, err := cfg.Build()
	if err != nil {
		return &ZapLogger{zap.NewNop().Sugar()}
	}
	return &ZapLogger{l.Sugar()}
}

func (l *ZapLogger) With(kv ...any) output.Logger {
	return &ZapLogger{l.SugaredLogger.With(kv...)}
}
func (l *ZapLogger) Named(name string) output.Logger {
	return &ZapLogger{l.SugaredLogger.Named(name)}
}

func (l *ZapLogger) Info(msg string, kv ...any)  { l.Infow(msg, kv...) }
func (l *ZapLogger) Warn(msg string, kv ...any)  { l.Warnw(msg, kv...) }
func (l *ZapLogger) Error(msg string, kv ...any) { l.Errorw(msg, kv...) }
func (l *ZapLogger) Debug(msg string, kv ...any) { l.Debugw(msg, kv...) }
func (l *ZapLogger) Sync() error                 { return l.SugaredLogger.Sync() }
