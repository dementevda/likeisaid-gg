package controllers

import (
	"io"
	"net/http"

	"github.com/dementevda/likeisaid-gg/backend/cmd/api/models"
	"github.com/dementevda/likeisaid-gg/backend/cmd/store"
)

func HandleAddUser(s *store.MongoStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{ID: 100, Password: "qweqwe", Email: "qweqwe"}
		user, err := s.AddUser(user)
		if err != nil {
			io.WriteString(w, err.Error())
		}
		io.WriteString(w, "ho")
	}
}
