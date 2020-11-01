package views

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/apimiddlewares"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"github.com/gorilla/mux"
)

// HandleTasks for get list of tasks or create task
func HandleTasks(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// POST
		if r.Method == http.MethodPost {

			decoder := json.NewDecoder(r.Body)
			newTaskJSON := &models.CreateTaskJson{}
			if err := decoder.Decode(newTaskJSON); err != nil {
				writeError(w, r, http.StatusBadRequest, err.Error(), "Wrong json")

				return
			}

			_, err := govalidator.ValidateStruct(newTaskJSON)
			if err != nil {
				writeError(w, r, http.StatusBadRequest, err.Error(), "Wrong parameters")
				return
			}

			if err := validDate(newTaskJSON.WaitBefore); err != nil {
				writeError(w, r, http.StatusBadRequest, err.Error(), "Wrong parameters")
				return
			}

			newTask := &models.CreateTask{
				CreateTaskJson: newTaskJSON,
				UserEmail:      r.Context().Value(apimiddlewares.CtxUserKey).(*models.User).Email,
				CreatedAt:      time.Now(),
			}

			_, err = s.AddTask(newTask)
			if err != nil {
				writeError(w, r, http.StatusInternalServerError, err.Error(), "I am broken")
				return
			}
			writeResponse(w, http.StatusAccepted, &emptyResponse{})
			return
		}

		// GET
		tasks, err := s.GetUserTasks(r.Context().Value(apimiddlewares.CtxUserKey).(*models.User).Email)
		if err != nil {
			writeError(w, r, http.StatusInternalServerError, err.Error(), "I am broken")
			return
		}

		writeResponse(w, http.StatusAccepted, tasks)
		return
	}
}

// HandleTask for get/update/delete current task
func HandleTask(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID := mux.Vars(r)["id"]

		task := checkTaskExists(s, taskID)
		if task == nil {
			writeError(w, r, http.StatusNotFound, "", "Task not found")
			return
		}
		if task.UserEmail != r.Context().Value(apimiddlewares.CtxUserKey).(*models.User).Email {
			writeError(w, r, http.StatusForbidden, "", "Permissions denied")
			return
		}

		// PATCH
		if r.Method == http.MethodPatch {

			decoder := json.NewDecoder(r.Body)
			updTaskJSON := &models.UpdateTaskJson{}

			if err := decoder.Decode(updTaskJSON); err != nil {
				writeError(w, r, http.StatusBadRequest, err.Error(), "Wrong json")
				return
			}

			updateTaskFields(task, updTaskJSON)

			if err := s.EditTask(task); err != nil {
				writeError(w, r, http.StatusInternalServerError, err.Error(), "I am broken")
				return
			}

			writeResponse(w, http.StatusAccepted, &emptyResponse{})
			return
		}

		// DELETE
		if r.Method == http.MethodDelete {
			s.DeleteTask(taskID)
			writeResponse(w, http.StatusAccepted, &emptyResponse{})
		}
	}
}

func checkTaskExists(s store.Store, id string) *models.Task {
	task, err := s.GetTaskByID(id)
	if err != nil {
		return nil
	}
	return task
}

func updateTaskFields(task *models.Task, upd *models.UpdateTaskJson) {
	if upd.Defendant != "" {
		task.Defendant = upd.Defendant
	}
	if upd.Title != "" {
		task.Title = upd.Title
	}
	nullTime := time.Time{}
	if upd.WaitBefore != nullTime {
		task.WaitBefore = upd.WaitBefore
	}
	if upd.Description != "" {
		task.Description = upd.Description
	}
}
