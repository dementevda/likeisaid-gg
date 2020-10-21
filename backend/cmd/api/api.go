package api

// API struct with core methods
type API struct {
	config *Config
}

// New returns new api server
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start api server
func (api *API) Start() error {
	return nil
}
