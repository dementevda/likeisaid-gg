package store

// Config for storage
type Config struct {
	DatabaseURL    string `toml:"database_url"`
	DatabaseName   string `toml:"database_name"`
	DatabaseUser   string `toml:"user"`
	DatabasePasswd string `toml:"password"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
