package apimiddlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxRequestIDKeyType string

// СtxRequestIDKey ...
const СtxRequestIDKey ctxRequestIDKeyType = "RequestID"

// RequestID adds request id to evry request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), СtxRequestIDKey, id)))
	})
}
