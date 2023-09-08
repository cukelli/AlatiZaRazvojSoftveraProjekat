package config

import (
	_ "encoding/json"
)

// swagger:model Config
type Config struct {
	// ID of the config
	// in: string
	ID string `json:"id"`

	// Name of the config
	// in: string
	Name string `json:"name"`

	// List of entries of the config
	// in: map[string]string
	Entries map[string]string `json:"entries"` //atribut entries kao [kljuc] prima string,kao vrednost string

	// GroupID of the config
	// in: string
	GroupID string `json:"group_id"`

	// Version of the config
	// in: string
	Version string `json:"version"`

	//List of labels of the config
	//in: map[string]string
	Labels string `json:"labels"`

	// Idempotency key associated with the configuration
	// in: string
	IdempotencyKey string `json:"idempotency_key"`
}
