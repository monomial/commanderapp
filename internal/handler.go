package internal

import (
	"commander-app/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, fmt.Sprintf("Method not allowed: %s", r.Method), http.StatusMethodNotAllowed)
			return
		}

		var cmdReq models.CommandRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&cmdReq); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var resp models.CommandResponse
		switch cmdReq.Type {
		case "ping":
			result, err := cmdr.Ping(cmdReq.Payload)
			if err != nil {
				resp = models.CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = models.CommandResponse{Success: true, Data: result}
			}
		case "sysinfo":
			result, err := cmdr.GetSystemInfo()
			if err != nil {
				resp = models.CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = models.CommandResponse{Success: true, Data: result}
			}
		default:
			resp = models.CommandResponse{Success: false, Error: fmt.Sprintf("Invalid command type: %s", cmdReq.Type)}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
