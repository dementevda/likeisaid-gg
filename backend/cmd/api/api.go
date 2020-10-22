package api

import (
	"net/http"

	"github.com/dementevda/likeisaid-gg/backend/cmd/api/controllers"
	"github.com/dementevda/likeisaid-gg/backend/cmd/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API struct with core methods
type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New returns new api server
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start api server
func (api *API) Start() error {
	if err := api.configureLogger(); err != nil {
		return err
	}

	api.configureRouter()

	if err := api.configureStore(); err != nil {
		return err
	}

	api.logger.Info("Starting API")

	return http.ListenAndServe(api.config.Port, api.router)
}

func (api *API) configureLogger() error {
	level, err := logrus.ParseLevel(api.config.LogLevel)
	if err != nil {
		return err
	}

	api.logger.SetLevel(level)

	return nil
}

func (api *API) configureRouter() {
	api.router.HandleFunc("/hello", controllers.HandleHello())
}

func (api *API) configureStore() error {
	store := store.New(api.config.Store)
	if err := store.Open(); err != nil {
		return err
	}

	api.store = store

	return nil
}
