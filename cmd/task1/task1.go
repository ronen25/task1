package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Register main hello world handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello-world")
	})

	// Get server port from environment
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatalf("Error: Missing `SERVER_PORT` environment variable")
	}

	log.Printf("Starting server on port .")

	bindAddr := fmt.Sprintf(":%s", serverPort)
	if err := http.ListenAndServe(bindAddr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
