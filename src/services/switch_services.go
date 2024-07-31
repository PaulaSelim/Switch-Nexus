package services

import "Switch-Nexus/src/utils"

type SwitchService struct {
	sshConnector *utils.SSHConnector
}

func NewSwitchService(sshConnector *utils.SSHConnector) *SwitchService {
	return &SwitchService{sshConnector: sshConnector}
}

func (s *SwitchService) GetCDPNeighbors() (string, error) {
	command := "show cdp neighbors"
	return s.sshConnector.RunCommand(command)
}
