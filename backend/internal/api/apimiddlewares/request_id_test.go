package apimiddlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestIDHeader(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handlerToTest := RequestID(nextHandler)

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	handlerToTest.ServeHTTP(recorder, request)
	assert.NotEqual(t, recorder.Header().Get("X-Request-ID"), "")
}
