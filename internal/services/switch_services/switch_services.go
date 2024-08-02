package switch_services

import "Switch-Nexus/internal/utilities/ssh"

type SwitchService struct {
	sshConnector *ssh.SSHConnector
}

func New(sshConnector *ssh.SSHConnector) *SwitchService {
	return &SwitchService{sshConnector: sshConnector}
}

func (s *SwitchService) GetCDPNeighbors() (string, error) {
	command := "show cdp neighbors"
	return s.sshConnector.RunCommand(command)
}
