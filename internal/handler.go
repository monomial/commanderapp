package internal

import (
	"commander-app/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received /execute request")

		if r.Method != http.MethodPost {
			http.Error(w, fmt.Sprintf("Method not allowed: %s", r.Method), http.StatusMethodNotAllowed)
			return
		}

		var cmdReq models.CommandRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&cmdReq); err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		log.Printf("Handling command of type: %s", cmdReq.Type)

		var resp models.CommandResponse
		switch cmdReq.Type {
		case "ping":
			result, err := cmdr.Ping(cmdReq.Payload)
			if err != nil {
				log.Printf("Ping command failed: %v", err)
				resp = models.CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = models.CommandResponse{Success: true, Data: result}
			}
		case "sysinfo":
			result, err := cmdr.GetSystemInfo()
			if err != nil {
				log.Printf("System info retrieval failed: %v", err)
				resp = models.CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = models.CommandResponse{Success: true, Data: result}
			}
		default:
			log.Printf("Unknown command type: %s", cmdReq.Type)
			resp = models.CommandResponse{Success: false, Error: fmt.Sprintf("Unknown command type: %s", cmdReq.Type)}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	}
}
