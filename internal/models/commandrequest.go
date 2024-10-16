package models

type CommandRequest struct {
	Type    string `json:"type"`    // "ping" or "sysinfo"
	Payload string `json:"payload"` // For ping, this is the host
}