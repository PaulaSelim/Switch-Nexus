package services

import "Switch-Nexus/src/utils"

type SwitchService struct {
	sshConnector *utils.SSHConnector
}

func NewSwitchService(sshConnector *utils.SSHConnector) *SwitchService {
	return &SwitchService{sshConnector: sshConnector}
}

func (s *SwitchService) ListDirectory(path string) (string, error) {
	command := "ls " + path
	return s.sshConnector.RunCommand(command)
}
