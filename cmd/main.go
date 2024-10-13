package main

import (
	"commander-app/internal"
	"log"
	"net/http"
)

func main() {
	commander := internal.NewCommander()
	server := &http.Server{
		Addr:    ":8080",
		Handler: internal.HandleRequests(commander),
	}
	log.Println("Server is running on port 8080")
	log.Fatal(server.ListenAndServe())
}
