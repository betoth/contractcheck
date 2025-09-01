package main

import (
	"embed"
	"log"

	applog "github.com/betoth/contractcheck/internal/adapter/logger"
	"github.com/betoth/contractcheck/internal/adapter/ui/wailsapp"
	"github.com/wailsapp/wails/v2"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	l := applog.New()
	defer l.Sync()

	if err := wails.Run(wailsapp.UIOptions(assets, l)); err != nil {
		log.Fatal(err)
	}
}
