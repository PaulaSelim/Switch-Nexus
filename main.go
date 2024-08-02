package main

import (
	"embed"
	"log"

	"Switch-Nexus/internal/config"
	"Switch-Nexus/internal/interfaces/wails_interfaces"
	"Switch-Nexus/internal/services/switch_services"
	"Switch-Nexus/internal/utilities/ssh"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	conf := config.LoadConfig()

	sshConnector, err := ssh.NewConnector(conf.SSHHost, conf.SSHUser, conf.SSHPassword)
	if err != nil {
		log.Fatalf("Failed to connect to SSH: %v", err)
		return
	}

	service := switch_services.New(sshConnector)
	app := wails_interfaces.New(service)

	err = wails.Run(&options.App{
		Title:  "Switch-Nexus",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 33, G: 33, B: 33, A: 1},
		// OnStartup:        app.startup,
		// OnShutdown:       app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatalf("Failed to run Wails app: %v", err)
	}
}
