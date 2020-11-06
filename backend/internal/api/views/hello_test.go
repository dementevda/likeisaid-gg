package views

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIHello(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	HandleHello().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "Hello, World!")
}
