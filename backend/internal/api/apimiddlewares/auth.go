package apimiddlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/apierrors"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

// CtxUserKey ...
const CtxUserKey ctxUserKeyType = "user"

type ctxUserKeyType string

// AuthUser дополнительное замыкание чтобы передать storage
func AuthUser(s store.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO get token and inspect
			user, err := s.GetUserByEmail("like@saifgd.gg")
			if err != nil {
				handleAuthErrors(err, w)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxUserKey, user)))
		})
	}
}

func handleAuthErrors(err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&apierrors.UserError{Message: err.Error(), ErrType: "Need registration"})
		fmt.Fprintln(w)
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&apierrors.UserError{Message: err.Error(), ErrType: "Error while searching user in middleware"})
		fmt.Fprintln(w)
		return
	}
}
