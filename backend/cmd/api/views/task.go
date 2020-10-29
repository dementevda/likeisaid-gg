package views

import (
	"fmt"
	"net/http"

	"github.com/dementevda/likeisaid-gg/backend/cmd/store"
	"github.com/gorilla/mux"
)

// HandleTasks for get list of tasks or create task
func HandleTasks(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO add task on post, selet on get

	}
}

// HandleTask for get/update/delete current task
func HandleTask(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO edit current task
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ID: %v\n", vars["id"])
	}
}
