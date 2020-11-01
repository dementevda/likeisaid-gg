package api

import (
	"net/http"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/apimiddlewares"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/views"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"github.com/dementevda/likeisaid-gg/backend/internal/store/mongostorage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API struct with core methods
type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
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
	// root router
	api.router.HandleFunc("/hello", views.HandleHello()).Methods("GET")

	// /users router
	userRouter := api.router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", views.HandleUsers(api.store)).Methods("POST")
	userRouter.HandleFunc("/{login}", views.HandleUser(api.store)).Methods("GET")

	// /tasks router
	taskRouter := api.router.PathPrefix("/tasks").Subrouter()
	taskRouter.Use(apimiddlewares.AuthUser(api.store))
	taskRouter.HandleFunc("", views.HandleTasks(api.store)).Methods("POST", "GET")
	taskRouter.HandleFunc("/{id}", views.HandleTask(api.store)).Methods("PATCH", "DELETE")
}

func (api *API) configureStore() error {
	api.logger.Info("Connecting to storage")
	store := mongostorage.New(api.config.Store)
	if err := store.Open(); err != nil {
		api.logger.Error("Fails to connect to storage")
		return err
	}
	api.logger.Info("Connected")
	api.store = store

	return nil
}
