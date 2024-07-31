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

func (a *WailsApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsApp) GetCDPNeighbors() (string, error) {
	return a.service.GetCDPNeighbors()
}
