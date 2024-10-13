package internal

import (
	"encoding/json"
	"net/http"
)

type CommandRequest struct {
	Type    string `json:"type"`    // "ping" or "sysinfo"
	Payload string `json:"payload"` // For ping, this is the host
}

type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func HandleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var cmdReq CommandRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&cmdReq); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var resp CommandResponse
		switch cmdReq.Type {
		case "ping":
			result, err := cmdr.Ping(cmdReq.Payload)
			if err != nil {
				resp = CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: result}
			}
		case "sysinfo":
			result, err := cmdr.GetSystemInfo()
			if err != nil {
				resp = CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: result}
			}
		default:
			resp = CommandResponse{Success: false, Error: "Invalid command type"}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
