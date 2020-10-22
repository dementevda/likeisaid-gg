package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIHello(t *testing.T) {
	api := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	api.handleHello().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "Hello, World!")
}
