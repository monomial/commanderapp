package internal

import (
	"errors"
	"net"
	"os"
	"os/exec"
	"time"
)

type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
	Successful bool          `json:"successful"`
	Time       time.Duration `json:"time"`
}

type SystemInfo struct {
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ip_address"`
}

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) Ping(host string) (PingResult, error) {
	start := time.Now()
	err := exec.Command("ping", "-c", "1", host).Run()
	if err != nil {
		return PingResult{Successful: false}, err
	}
	duration := time.Since(start)
	return PingResult{Successful: true, Time: duration}, nil
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return SystemInfo{}, err
	}

	var ipAddr string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ipAddr = ipNet.IP.String()
			break
		}
	}

	if ipAddr == "" {
		return SystemInfo{}, errors.New("could not determine IP address")
	}

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: ipAddr,
	}, nil
}
