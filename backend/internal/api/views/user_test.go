package views

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/apimiddlewares"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"github.com/dementevda/likeisaid-gg/backend/internal/store/teststorage"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func initStore() store.Store {
	return &teststorage.TestStorage{
		Users: make(map[string]*models.User),
		Tasks: make(map[string]*models.Task),
	}
}

func addUser(s store.Store) *models.User {
	user := &models.CreateUser{
		Email: "test@test.gg",
		Login: "test",
	}
	u, _ := s.AddUser(user)
	return u
}

type testCaseUserPOST struct {
	dataJSON   string
	statusCode int
	name       string
}

func TestHandleUsers(t *testing.T) {
	testCases := []testCaseUserPOST{
		testCaseUserPOST{
			dataJSON: `{
			"login": "test",
			"email": "test@test.gg"
			}`,
			statusCode: http.StatusAccepted,
			name:       "OK",
		},
		testCaseUserPOST{
			dataJSON: `{
			"login": "test",,
			"email": "test@test.gg"
			}`,
			statusCode: http.StatusBadRequest,
			name:       "Bad JSON",
		},
		testCaseUserPOST{
			dataJSON: `{
			"login": "test",
			"email": "testtest.gg"
			}`,
			statusCode: http.StatusBadRequest,
			name:       "Bad email",
		},
		testCaseUserPOST{
			dataJSON: `{
			"login": "test",
			"email": "test@test.gg"
			}`,
			statusCode: http.StatusBadRequest,
			name:       "User exists",
		},
		testCaseUserPOST{
			dataJSON: `{
			"login": "internalerror",
			"email": "test@test.gg"
			}`,
			statusCode: http.StatusInternalServerError,
			name:       "Internal error",
		},
	}

	store := initStore()
	handler := HandleUsers(store)
	for _, testCase := range testCases {
		request := httptest.NewRequest("POST", "/", strings.NewReader(testCase.dataJSON))
		request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)
		assert.Equal(t, testCase.statusCode, recorder.Code, testCase.name)
	}
}

func TestHandleUser(t *testing.T) {
	store := initStore()
	_ = addUser(store)

	router := func() *mux.Router {
		r := mux.NewRouter()
		r.HandleFunc("/users/{login}", HandleUser(store))
		return r
	}()

	request := httptest.NewRequest("GET", "/users/test", nil)
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, 200)
}

func TestHandleUserNotFound(t *testing.T) {
	store := initStore()

	router := func() *mux.Router {
		r := mux.NewRouter()
		r.HandleFunc("/users/{login}", HandleUser(store))
		return r
	}()

	request := httptest.NewRequest("GET", "/users/test", nil)
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, 404)
}

func TestHandleUserInternal(t *testing.T) {
	store := initStore()
	// _ = addUser(store)

	router := func() *mux.Router {
		r := mux.NewRouter()
		r.HandleFunc("/users/{login}", HandleUser(store))
		return r
	}()

	request := httptest.NewRequest("GET", "/users/internalerror", nil)
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, 500)
}
