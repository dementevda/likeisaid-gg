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
	store  *store.MongoStorage
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

	if err := api.configureStore(); err != nil {
		return err
	}
	api.configureRouter()

	api.logger.Info("Starting API")

	return http.ListenAndServe(api.config.Port, api.router)
}

// Stop ...
func (api *API) Stop() {
	api.store.Close()
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
	api.router.HandleFunc("/user", controllers.HandleAddUser(api.store)).Methods("POST")
}

func (api *API) configureStore() error {
	store := store.New(api.config.Store)
	if err := store.Open(); err != nil {
		return err
	}

	api.store = store

	return nil
}
