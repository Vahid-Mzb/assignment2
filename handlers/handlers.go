package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"kvstore/models"
	"kvstore/store"

	"github.com/gorilla/mux"
)

type Handler struct {
	store *store.Store
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{store: s}
}

// PutHandler handles PUT /objects requests
func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req models.PutRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.Key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}

	// Store the data
	h.store.Put(req.Key, req.Value)

	// Persist to disk
	if err := h.store.Save(); err != nil {
		log.Printf("Error saving data: %v\n", err)
		http.Error(w, "Failed to persist data", http.StatusInternalServerError)
		return
	}

	log.Printf("Stored key: %s\n", req.Key)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}

	value, ok := h.store.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(value)
	log.Printf("Retrieved key: %s\n", key)
}
