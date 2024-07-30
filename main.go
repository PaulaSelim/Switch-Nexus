package main

import (
	"embed"

	"Switch-Nexus/src/config"
	"Switch-Nexus/src/interfaces"
	"Switch-Nexus/src/middleware"
	"Switch-Nexus/src/services"
	"Switch-Nexus/src/utils"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	conf := config.LoadConfig()

	sshConnector, err := utils.NewSSHConnector(conf.SSHHost, conf.SSHUser, conf.SSHPassword)
	if err != nil {
		middleware.LogError(&middleware.AppError{
			Code:    middleware.ErrCodeConnectionFailed,
			Message: "SSH connection failed",
		})
		return
	}

	service := services.NewSwitchService(sshConnector)
	app := interfaces.NewWailsApp(service)

	err = wails.Run(&options.App{
		Title:  "Switch-Nexus",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		middleware.LogError(&middleware.AppError{
			Code:    middleware.ErrCodeConnectionFailed,
			Message: err.Error(),
		})
	}
}
