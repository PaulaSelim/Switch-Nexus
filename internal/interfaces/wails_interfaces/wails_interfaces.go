package wails_interfaces

import (
	"Switch-Nexus/internal/utilities/ping"
	"Switch-Nexus/internal/utilities/ssh"
	"context"
	"log"
)

type WailsInterface struct {
	ctx  context.Context
	user string
	pass string
}

func New(user, password string) *WailsInterface {
	return &WailsInterface{user: user, pass: password}
}

func (a *WailsInterface) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsInterface) Shutdown(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsInterface) GetStatus(host string) (string, error) {
	pingStatus, err := ping.PingServer(host)
	if err != nil {
		return "", err
	}
	return pingStatus, nil
}

func (a *WailsInterface) GetVLAN(host string) (string, error) {
	log.Printf("GetVLAN on: %s, with credentials: user(%s) pass(%s)", host, a.user, a.pass)
	sshConnector, err := ssh.NewConnector(host, a.user, a.pass)
	if err != nil {
		log.Printf("Failed to connect to SSH: %v", err)
		return "", err
	}
	// command := "ipconfig"
	command := "show vlan"
	return sshConnector.RunCommand(command)
}

func (a *WailsInterface) GetCDPNeighbors(host string) (string, error) {
	log.Printf("GetCDPNeighbors on: %s, with credentials: user(%s) pass(%s)", host, a.user, a.pass)
	sshConnector, err := ssh.NewConnector(host, a.user, a.pass)
	if err != nil {
		log.Printf("Failed to connect to SSH: %v", err)
		return "", err
	}
	// command := "ipconfig"
	command := "show cdp neighbors"
	return sshConnector.RunCommand(command)
}

func (a *WailsInterface) GetIPInterface(host string) (string, error) {
	log.Printf("GetIPInterface on: %s, with credentials: user(%s) pass(%s)", host, a.user, a.pass)
	sshConnector, err := ssh.NewConnector(host, a.user, a.pass)
	if err != nil {
		log.Printf("Failed to connect to SSH: %v", err)
		return "", err
	}
	// command := "systeminfo"
	command := "show ip interface brief"
	return sshConnector.RunCommand(command)
}
