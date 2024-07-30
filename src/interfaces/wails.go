package interfaces

import (
	"context"

	"Switch-Nexus/src/services"
)

type WailsApp struct {
	ctx     context.Context
	service *services.SwitchService
}

func NewWailsApp(service *services.SwitchService) *WailsApp {
	return &WailsApp{service: service}
}

func (a *WailsApp) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsApp) ListDirectory(path string) (string, error) {
	return a.service.ListDirectory(path)
}
