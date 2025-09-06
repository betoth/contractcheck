package main

import (
	"embed"
	"io/fs"
	"log"

	applog "github.com/betoth/contractcheck/internal/adapter/logger"
	"github.com/betoth/contractcheck/internal/adapter/ui/wailsapp"
	"github.com/wailsapp/wails/v2"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize backend logger (zap)
	l := applog.New()
	defer l.Sync()

	// Ensure embedded assets point to "frontend/dist"
	dist, err := fs.Sub(assets, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	// Run Wails with centralized options
	if err := wails.Run(wailsapp.UIOptions(dist, l)); err != nil {
		log.Fatal(err)
	}
}
