package ping

import (
	"fmt"

	"github.com/go-ping/ping"
)

func PingServer(host string) (string, error) {
	pinger, err := ping.NewPinger(host)
	pinger.SetPrivileged(true)
	if err != nil {
		fmt.Printf("Failed to create pinger for server: %v \n %v", host, err)
		return "", err
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		fmt.Printf("Failed to ping server: %v \n %v", host, err)
		return "", err
	}
	fmt.Printf("Ping succeeded for server: %v \n Ping statistics: %v\n", host, pinger.Statistics())
	return "success", nil
}
