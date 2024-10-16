package models

import "time"

type PingResult struct {
	Successful bool          `json:"successful"`
	Time       time.Duration `json:"time"`
}
