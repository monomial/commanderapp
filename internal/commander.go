package internal

import (
	"commander-app/internal/models"
	"errors"
	"net"
	"os"
	"os/exec"
	"time"
)

type Commander interface {
	Ping(host string) (models.PingResult, error)
	GetSystemInfo() (models.SystemInfo, error)
}

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) Ping(host string) (models.PingResult, error) {
	start := time.Now()
	err := exec.Command("ping", "-c", "1", host).Run()
	if err != nil {
		return models.PingResult{Successful: false}, err
	}
	duration := time.Since(start)
	return models.PingResult{Successful: true, Time: duration}, nil
}

func (c *commander) GetSystemInfo() (models.SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return models.SystemInfo{}, err
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return models.SystemInfo{}, err
	}

	var ipAddr string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ipAddr = ipNet.IP.String()
			break
		}
	}

	if ipAddr == "" {
		return models.SystemInfo{}, errors.New("could not determine IP address")
	}

	return models.SystemInfo{
		Hostname:  hostname,
		IPAddress: ipAddr,
	}, nil
}
