package views

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/apimiddlewares"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addTask(s store.Store, email string, id bool) *models.Task {
	taskJSON := &models.CreateTaskJson{
		Title:       "test title",
		Description: "test descrition",
		WaitBefore:  time.Now().Add(time.Hour * 24),
	}
	taskCreate := models.CreateTask{
		CreateTaskJson: taskJSON,
		UserEmail:      email,
	}

	t, _ := s.AddTask(&taskCreate)
	if id {
		t.ID = primitive.NewObjectID()
	}

	return t
}

type testCaseTaskPOST struct {
	dataJSON   string
	statusCode int
	name       string
}

func TestHandleTasksPOST(t *testing.T) {
	testCases := []testCaseTaskPOST{
		testCaseTaskPOST{
			dataJSON: `{
			"wait_before": "2021-12-20T14:00:13Z",
			"description": "test description",
			"title": "test title"
		}`,
			statusCode: http.StatusAccepted,
			name:       "OK",
		},
		testCaseTaskPOST{
			dataJSON: `{
			"wait_before": "2021-12-20T14:00:13Z",
			"description": "test description",
			"title": "internalerror"
			}
			`,
			statusCode: http.StatusInternalServerError,
			name:       "Internal error",
		},
		testCaseTaskPOST{
			dataJSON: `{
			"wait_before": "2021-12-202T14:00:13Z",
			"description": "test description",
			"title": "test title"
			}
			`,
			statusCode: http.StatusBadRequest,
			name:       "Bad Date",
		},
		testCaseTaskPOST{
			dataJSON: `{
			"wait_before": "2021-12-20T14:00:13Z",
			"description": "test description",,
			"title": "test title"
			}
			`,
			statusCode: http.StatusBadRequest,
			name:       "Bad JSON",
		},
		testCaseTaskPOST{
			dataJSON: `{
			"wait_before": "2021-12-20T14:00:13Z",
			"description": "test description"
			}
			`,
			statusCode: http.StatusBadRequest,
			name:       "Bad params",
		},
	}

	for _, testCase := range testCases {
		store := initStore()
		user := addUser(store)

		router := func() *mux.Router {
			r := mux.NewRouter()
			r.HandleFunc("/tasks", HandleTasks(store))
			return r
		}()
		request := httptest.NewRequest("POST", "/tasks", strings.NewReader(testCase.dataJSON))
		request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
		request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.CtxUserKey, user))

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		assert.Equal(t, testCase.statusCode, recorder.Code, testCase.name)
	}
}

func TestHandleTasksGET(t *testing.T) {
	store := initStore()
	user := addUser(store)

	router := func() *mux.Router {
		r := mux.NewRouter()
		r.HandleFunc("/tasks", HandleTasks(store))
		return r
	}()

	request := httptest.NewRequest("GET", "/tasks", nil)
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.CtxUserKey, user))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
}

func TestHandleTasksGETInternal(t *testing.T) {
	store := initStore()
	user := addUser(store)
	user.Email = "internalerror"

	router := func() *mux.Router {
		r := mux.NewRouter()
		r.HandleFunc("/tasks", HandleTasks(store))
		return r
	}()

	request := httptest.NewRequest("GET", "/tasks", nil)
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
	request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.CtxUserKey, user))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	assert.Equal(t, 500, recorder.Code)
}

type testCaseHandleTask struct {
	name       string
	dataJSON   string
	statusCode int
	method     string
	email      string
	taskID     bool
}

func TestHandleTask(t *testing.T) {
	testCases := []testCaseHandleTask{
		testCaseHandleTask{
			name:       "Task not exists",
			statusCode: http.StatusNotFound,
			method:     http.MethodGet,
			email:      "test@test.gg",
			taskID:     true,
			dataJSON:   `{}`,
		},
		testCaseHandleTask{
			name:       "Wrong task",
			statusCode: http.StatusForbidden,
			method:     http.MethodGet,
			email:      "another@test.gg",
			taskID:     false,
			dataJSON:   `{}`,
		},
		testCaseHandleTask{
			name:       "PATCH bad json",
			statusCode: http.StatusBadRequest,
			method:     http.MethodPatch,
			email:      "test@test.gg",
			taskID:     false,
			dataJSON:   `{,}`,
		},
		testCaseHandleTask{
			name:       "PATCH bad params",
			statusCode: http.StatusInternalServerError,
			method:     http.MethodPatch,
			email:      "test@test.gg",
			taskID:     false,
			dataJSON:   `{"title": "internalerror"}`,
		},
		testCaseHandleTask{
			name:       "PATCH OK",
			statusCode: http.StatusAccepted,
			method:     http.MethodPatch,
			email:      "test@test.gg",
			taskID:     false,
			dataJSON: `{
				"title": "title",
				"description": "description",
				"wait_before": "2333-12-15T14:00:13Z"
			}`,
		},
		testCaseHandleTask{
			name:       "DELETE OK",
			statusCode: http.StatusAccepted,
			method:     http.MethodDelete,
			email:      "test@test.gg",
			taskID:     false,
			dataJSON:   ``,
		},
	}

	for _, testCase := range testCases {
		store := initStore()
		user := addUser(store)
		router := func() *mux.Router {
			r := mux.NewRouter()
			r.HandleFunc("/tasks/{id}", HandleTask(store))
			return r
		}()

		task := addTask(store, testCase.email, testCase.taskID)

		request := httptest.NewRequest(testCase.method, "/tasks/"+task.ID.String(), strings.NewReader(testCase.dataJSON))
		request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.СtxRequestIDKey, "test"))
		request = request.WithContext(context.WithValue(request.Context(), apimiddlewares.CtxUserKey, user))

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		assert.Equal(t, testCase.statusCode, recorder.Code, testCase.name)
	}
}
