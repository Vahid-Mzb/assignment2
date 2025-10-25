package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"kvstore/handlers"
	"kvstore/store"

	"github.com/gorilla/mux"
)

const port = "8080"

func main() {
	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "data/store.json"
	}

	log.Printf("Using data path: %s\n", dataPath)

	// Initialize store
	kvStore, err := store.NewStore(dataPath)
	if err != nil {
		log.Fatalf("Failed to initialize store: %v\n", err)
	}

	// Load existing data
	if err := kvStore.Load(); err != nil {
		log.Printf("Warning: failed to load existing data: %v\n", err)
	}

	// Initialize handlers
	h := handlers.NewHandler(kvStore)

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/objects", h.PutHandler).Methods("PUT")
	r.HandleFunc("/objects/{key}", h.GetHandler).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
