package api

import (
	"io"
	"net/http"

	"github.com/dementevda/likeisaid-gg/backend/cmd/api/models"
	"github.com/dementevda/likeisaid-gg/backend/cmd/store"
)

func (api *API) handleAddUser(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{Password: "qweqweqweqweqweqwe", Email: "qweqwe"}
		user, err := s.AddUser(user)
		if err != nil {
			io.WriteString(w, err.Error())
		}
		io.WriteString(w, "Ok!")
	}
}

func (api *API) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!")
	}
}
