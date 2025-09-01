package wailsapp

import (
	"context"
	"io/fs"

	"github.com/betoth/contractcheck/internal/application/ports/output"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// App is the UI adapter bound to the Wails runtime.
// Keep domain logic out of here; call use cases via ports instead.
type App struct {
	ctx context.Context
	log output.Logger
}

// New constructs the UI adapter. If your logger supports Named(), tag logs with "ui".
func New(log output.Logger) *App {
	if log != nil {
		log = log.Named("ui")
	}
	return &App{log: log}
}

// Startup is called when the Wails runtime is ready to start the app.
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

// Shutdown is called when the application is about to quit.
func (a *App) Shutdown(ctx context.Context) {
	if a.log != nil {
		a.log.Info("UI shutdown")
		_ = a.log.Sync()
	}
}

// UIOptions builds the Wails options using the provided embedded asset FS.
// The embed stays in main to keep paths stable.
func UIOptions(assets fs.FS, log output.Logger) *options.App {
	app := New(log)
	return &options.App{
		Title:  "ContractCheck",
		Width:  1200,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.Startup,
		OnDomReady: app.DomReady,
		OnShutdown: app.Shutdown,
		Bind: []interface{}{
			app, // exported methods become available to the frontend
		},
	}
}
