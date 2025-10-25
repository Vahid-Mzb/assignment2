package models

import "encoding/json"

type PutRequest struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}
