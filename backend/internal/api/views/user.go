package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/apierrors"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

var decoder = schema.NewDecoder()

// HandleUsers adds users to database on POST query
func HandleUsers(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var newUser *models.CreateUser = &models.CreateUser{}

		if err := decoder.Decode(newUser); err != nil {
			handleError(w, http.StatusBadRequest, &apierrors.UserError{Message: err.Error(), ErrType: "JSON decode error"})
			return
		}

		_, err := govalidator.ValidateStruct(newUser)
		if err != nil {
			handleError(w, http.StatusBadRequest, &apierrors.UserError{Message: err.Error(), ErrType: "Wrong parameters"})
			return
		}

		user, err := s.AddUser(newUser)
		switch {
		case isDup(err):
			handleError(w, http.StatusBadRequest, &apierrors.UserError{Message: "User alredy in db", ErrType: "Exists"})
			return
		case err != nil:
			handleError(w, http.StatusInternalServerError, &apierrors.UserError{Message: err.Error(), ErrType: "Error in saving user"})
			return
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(user)
		fmt.Fprintln(w)
		return
	}
}

// HandleUser get user by login
func HandleUser(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := mux.Vars(r)["login"]

		user, err := s.FindUser(login)
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			handleError(w, http.StatusNotFound, &apierrors.UserError{Message: err.Error(), ErrType: "Not Found"})
			return
		case err != nil:
			handleError(w, http.StatusInternalServerError, &apierrors.UserError{Message: err.Error(), ErrType: "Error while searching user"})
			return
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(user)
		fmt.Fprintln(w)
		return
	}
}
