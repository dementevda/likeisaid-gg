package apimiddlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// RequestTimer adds request id to evry request
func RequestTimer(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			log := logger.WithFields(logrus.Fields{
				"method":      r.Method,
				"URI":         r.RequestURI,
				"remote_addr": r.RemoteAddr,
				"request_id":  r.Context().Value(СtxRequestIDKey).(string),
				"duration":    time.Now().Sub(start),
			})
			log.Infof("[%v]", time.Now())
		})
	}
}
