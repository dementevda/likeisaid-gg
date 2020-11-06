package apimiddlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestRequestTimer(t *testing.T) {
	// log := logrus.New()
	log, hook := test.NewNullLogger()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	requestTimer := RequestTimer(log)
	handlerToTest := requestTimer(nextHandler)

	request := httptest.NewRequest("GET", "/", nil)
	request = request.WithContext(context.WithValue(request.Context(), СtxRequestIDKey, "test"))
	recorder := httptest.NewRecorder()

	handlerToTest.ServeHTTP(recorder, request)
	assert.Equal(t, hook.LastEntry().Data["request_id"], "test")
}
