package main

import (
	"embed"
	"log"

	"Switch-Nexus/internal/config"
	"Switch-Nexus/internal/interfaces/wails_interfaces"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	conf := config.LoadConfig()
	
	app := wails_interfaces.New(conf.SSHUser, conf.SSHPassword)

	err := wails.Run(&options.App{
		Title:  "Switch-Nexus",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatalf("Failed to run Wails app: %v", err)
	}
}
