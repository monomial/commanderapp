package internal

import (
	"commander-app/internal/models"
	"errors"
	"log"
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
	log.Printf("Executing ping for host: %s", host)

	start := time.Now()
	err := exec.Command("ping", "-c", "1", host).Run()
	if err != nil {
		log.Printf("Ping failed for host: %s, error: %v", host, err)
		return models.PingResult{Successful: false}, err
	}
	duration := time.Since(start)
	log.Printf("Ping successful for host: %s, time: %s", host, duration.String())
	return models.PingResult{Successful: true, Time: duration}, nil
}

func (c *commander) GetSystemInfo() (models.SystemInfo, error) {
	log.Println("Retrieving system information")

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Failed to get hostname: %v", err)
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
		log.Printf("could not determine IP address")
		return models.SystemInfo{}, errors.New("could not determine IP address")
	}

	log.Printf("System info retrieved: Hostname=%s, IPAddress=%s", hostname, ipAddr)

	return models.SystemInfo{
		Hostname:  hostname,
		IPAddress: ipAddr,
	}, nil
}
