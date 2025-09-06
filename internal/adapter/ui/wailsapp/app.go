package wailsapp

import (
	"context"
	"io/fs"

	"github.com/betoth/contractcheck/internal/application/ports/output"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// App is the UI adapter bound to the Wails runtime.
// It exposes minimal methods (e.g. Version) to the frontend.
// Keep domain/business logic outside of this layer.
type App struct {
	ctx context.Context
	log output.Logger
}

// New constructs the UI adapter.
// Logger is scoped with "ui" for filtering in zap logs.
func New(log output.Logger) *App {
	if log != nil {
		log = log.Named("ui")
	}
	return &App{log: log}
}

// Startup is called when the Wails runtime is ready.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	if a.log != nil {
		a.log.Info("UI startup")
	}
}

// DomReady is called when the DOM is fully loaded.
func (a *App) DomReady(ctx context.Context) {
	if a.log != nil {
		a.log.Info("UI DOM ready")
	}
}

// Shutdown is called when the application is quitting.
func (a *App) Shutdown(ctx context.Context) {
	if a.log != nil {
		a.log.Info("UI shutdown")
		_ = a.log.Sync()
	}
}

// Version exposes the app version to the frontend.
func (a *App) Version() string {
	// TODO: inject real version (ex: build flag, pkg/version)
	return "0.1.0"
}

// UIOptions builds the Wails app options, binding all frontend-facing APIs.
// This is the single entrypoint consumed by main.go.
func UIOptions(assets fs.FS, log output.Logger) *options.App {
	app := New(log)

	return &options.App{
		Title:            "ContractCheck",
		WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.Startup,
		OnDomReady: app.DomReady,
		OnShutdown: app.Shutdown,
		Bind: []interface{}{
			app,                  // Provides Version()
			NewLoggerBridge(log), // Provides frontend logging bridge
		},
	}
}
