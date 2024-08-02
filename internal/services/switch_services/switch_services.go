package switch_services

import "Switch-Nexus/internal/utilities/ssh"

type SwitchService struct {
	sshConnector *ssh.SSHConnector
}

func New(sshConnector *ssh.SSHConnector) *SwitchService {
	return &SwitchService{sshConnector: sshConnector}
}

func (s *SwitchService) GetStatus() (string, error) {
	command := "ipconfig"
	return s.sshConnector.RunCommand(command)
}

func (s *SwitchService) GetVLAN() (string, error) {
	command := "ipconfig"
	return s.sshConnector.RunCommand(command)
}

func (s *SwitchService) GetCDPNeighbors() (string, error) {
	command := "ipconfig"
	return s.sshConnector.RunCommand(command)
}

func (s *SwitchService) GetIPInterface() (string, error) {
	command := "getmac"
	return s.sshConnector.RunCommand(command)
}