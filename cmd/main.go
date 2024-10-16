package main

import (
	"commander-app/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting the application...")

	// Get the port from the environment or default to 8080
	port := os.Getenv("COMMANDERAPPPORT")
	if port == "" {
		port = "8080"
		log.Println("Defaulting port to " + port)
	}

	commander := internal.NewCommander()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: internal.HandleRequests(commander),
	}

	log.Printf("Server is running on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
