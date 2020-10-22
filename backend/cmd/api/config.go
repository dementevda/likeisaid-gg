package api

import "github.com/dementevda/likeisaid-gg/backend/cmd/store"

// Config for API
type Config struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig returns new Config struct
func NewConfig() *Config {
	return &Config{
		Port:     ":8000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
