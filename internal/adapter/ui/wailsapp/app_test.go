package wailsapp_test

import (
	"context"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/betoth/contractcheck/internal/adapter/ui/wailsapp"
	"github.com/betoth/contractcheck/internal/application/ports/output"
)

// stubLogger is a minimal test double for output.Logger.
// It records calls so we can assert lifecycle behavior.
// NOTE: Named/With return the same instance to keep a single sink of calls.
type stubLogger struct {
	name  string
	calls []string
}

func (l *stubLogger) With(kv ...any) output.Logger {
	// Keep it simple for unit tests: ignore fields and return same logger.
	return l
}

func (l *stubLogger) Named(name string) output.Logger {
	l.name = name
	return l
}

func (l *stubLogger) Info(msg string, kv ...any)  { l.calls = append(l.calls, "INFO:"+msg) }
func (l *stubLogger) Warn(msg string, kv ...any)  { l.calls = append(l.calls, "WARN:"+msg) }
func (l *stubLogger) Error(msg string, kv ...any) { l.calls = append(l.calls, "ERROR:"+msg) }
func (l *stubLogger) Debug(msg string, kv ...any) { l.calls = append(l.calls, "DEBUG:"+msg) }
func (l *stubLogger) Sync() error                 { return nil }

// mkAssets builds an in-memory fs.FS that simulates a built frontend.
// Wails expects "frontend/dist/index.html" to exist at runtime when embedding.
func mkAssets() fs.FS {
	return fstest.MapFS{
		"frontend/dist/index.html": &fstest.MapFile{
			Data: []byte("<!doctype html><html><head><meta charset=\"utf-8\"></head><body><div id='app'></div></body></html>"),
		},
	}
}

func TestUIOptions_BootLifecycle_LogsAndBinds(t *testing.T) {
	assets := mkAssets()
	log := &stubLogger{}

	opts := wailsapp.UIOptions(assets, log)
	if opts == nil {
		t.Fatal("expected options.App, got nil")
	}
	if opts.Title != "ContractCheck" {
		t.Fatalf("expected Title=ContractCheck, got %q", opts.Title)
	}
	if opts.AssetServer == nil || opts.AssetServer.Assets == nil {
		t.Fatal("expected AssetServer with non-nil Assets")
	}
	if len(opts.Bind) != 1 {
		t.Fatalf("expected exactly 1 bound object, got %d", len(opts.Bind))
	}
	// Ensure the bound object is of type *wailsapp.App
	if _, ok := opts.Bind[0].(*wailsapp.App); !ok {
		t.Fatalf("expected Bind[0] to be *wailsapp.App, got %T", opts.Bind[0])
	}

	// Exercise lifecycle handlers (should not panic and should log)
	ctx := context.Background()
	opts.OnStartup(ctx)
	opts.OnDomReady(ctx)
	opts.OnShutdown(ctx)

	// Assert our stub captured the expected lifecycle logs
	want := []string{"INFO:UI startup", "INFO:UI DOM ready", "INFO:UI shutdown"}
	have := log.calls
	if len(have) < len(want) {
		t.Fatalf("expected at least %d log entries, got %d: %#v", len(want), len(have), have)
	}
	for _, w := range want {
		found := false
		for _, h := range have {
			if h == w {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing log entry %q in calls: %#v", w, have)
		}
	}
}

func TestUIOptions_NoLogger_NoPanic(t *testing.T) {
	assets := mkAssets()
	opts := wailsapp.UIOptions(assets, nil)
	if opts == nil {
		t.Fatal("expected options.App, got nil")
	}
	// Lifecycle should be safe without a logger
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("unexpected panic without logger: %v", r)
		}
	}()
	ctx := context.Background()
	opts.OnStartup(ctx)
	opts.OnDomReady(ctx)
	opts.OnShutdown(ctx)
}
