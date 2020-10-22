package store

// Config for storage
type Config struct {
	DatabaseURL  string `toml:"database_url"`
	DatabaseName string
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
