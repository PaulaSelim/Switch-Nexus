package wails_interfaces

import (
	"context"

	"Switch-Nexus/internal/services/switch_services"
)

type WailsInterface struct {
	ctx     context.Context
	service *switch_services.SwitchService
}

func New(service *switch_services.SwitchService) *WailsInterface {
	return &WailsInterface{service: service}
}

func (a *WailsInterface) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsInterface) Shutdown(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsInterface) GetStatus() (string, error) {
	return a.service.GetStatus()
}

func (a *WailsInterface) GetVLAN() (string, error) {
	return a.service.GetVLAN()
}

func (a *WailsInterface) GetCDPNeighbors() (string, error) {
	return a.service.GetCDPNeighbors()
}

func (a *WailsInterface) GetIPInterface() (string, error) {
	return a.service.GetIPInterface()
}
