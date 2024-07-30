package utils

import (
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"
)

type SSHConnector struct {
	client *ssh.Client
}

func NewSSHConnector(host, user, password string) (*SSHConnector, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to SSH server")
		return nil, err
	}

	return &SSHConnector{client: client}, nil
}

func (s *SSHConnector) RunCommand(command string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
