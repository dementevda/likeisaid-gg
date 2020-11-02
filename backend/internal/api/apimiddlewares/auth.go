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

type ctxUserKeyType string

// CtxUserKey ...
const CtxUserKey ctxUserKeyType = "user"

// AuthUser дополнительное замыкание чтобы передать storage
func AuthUser(s store.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO get token and inspect
			user, err := s.GetUserByEmail("like@said.gg")
			if err != nil {
				handleAuthErrors(err, w, r)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxUserKey, user)))
		})
	}
}

func handleAuthErrors(err error, w http.ResponseWriter, r *http.Request) {
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&apierrors.APIError{
			Message:   err.Error(),
			ErrType:   "Need registration",
			RequestID: r.Context().Value(СtxRequestIDKey).(string)})
		fmt.Fprintln(w)
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&apierrors.APIError{
			Message:   err.Error(),
			ErrType:   "Error while searching user in middleware",
			RequestID: r.Context().Value(СtxRequestIDKey).(string)})
		fmt.Fprintln(w)
		return
	}
}
