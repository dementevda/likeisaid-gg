package api

// Config for API
type Config struct {
	Port string
}

// NewConfig returns new Config struct
func NewConfig() *Config {
	return &Config{
		Port: ":8000",
	}
}
