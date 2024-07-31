package utils

import (
	"log"
	"time"

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
		log.Fatalf("Failed to dial: %s", err)
		return nil, err
	}

	return &SSHConnector{client: client}, nil
}

func (s *SSHConnector) RunCommand(command string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		log.Printf("Failed to create session: %s", err)
		return "", err
	}
	defer session.Close()

	log.Printf("Running command: %s", command)

	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Printf("Failed to run command: %s", err)
		return "", err
	}

	log.Printf("Output: %s", output)

	return string(output), nil
}
