package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Store struct {
	mu       sync.RWMutex
	data     map[string]json.RawMessage
	filePath string
}

func NewStore(filePath string) (*Store, error) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	return &Store{
		data:     make(map[string]json.RawMessage),
		filePath: filePath,
	}, nil
}

// Load reads data from the file into memory
func (s *Store) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("No existing data file found, starting fresh")
			return nil
		}
		return fmt.Errorf("failed to open data file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat data file: %w", err)
	}
	if stat.Size() == 0 {
		log.Println("Data file is empty, starting fresh")
		return nil
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.data); err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	log.Printf("Loaded %d keys from storage\n", len(s.data))
	return nil
}

// Save writes the current data to the file
func (s *Store) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.Create(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to create data file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.data); err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	return nil
}

// Get retrieves a value by key
func (s *Store) Get(key string) (json.RawMessage, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// Put stores a key-value pair
func (s *Store) Put(key string, value json.RawMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Count returns the number of keys in the store
func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}
