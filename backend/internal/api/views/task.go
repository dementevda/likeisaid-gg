package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/apierrors"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/apimiddlewares"
	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"github.com/gorilla/mux"
)

// HandleTasks for get list of tasks or create task
func HandleTasks(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			decoder := json.NewDecoder(r.Body)
			newTaskJSON := &models.CreateTaskJson{}
			if err := decoder.Decode(newTaskJSON); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&apierrors.TaskError{Message: err.Error(), ErrType: "Wrong json"})
				fmt.Fprintln(w)
				return
			}

			_, err := govalidator.ValidateStruct(newTaskJSON)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&apierrors.TaskError{Message: err.Error(), ErrType: "Wrong parameters"})
				fmt.Fprintln(w)
				return
			}

			if err := validDate(newTaskJSON.WaitBefore); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&apierrors.TaskError{Message: err.Error(), ErrType: "Wrong parameters"})
				fmt.Fprintln(w)
				return
			}

			newTask := &models.CreateTask{
				CreateTaskJson: newTaskJSON,
				UserEmail:      r.Context().Value(apimiddlewares.CtxUserKey).(*models.User).Email,
				CreatedAt:      time.Now(),
			}

			user, err := s.AddTask(newTask)

			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(user)
			fmt.Fprintln(w)
			return
		}

		// GET
		tasks, err := s.GetUserTasks(r.Context().Value(apimiddlewares.CtxUserKey).(*models.User).Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&apierrors.TaskError{Message: err.Error(), ErrType: "I am broken"})
			fmt.Fprintln(w)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(tasks)
		fmt.Fprintln(w)
		return
	}
}

// HandleTask for get/update/delete current task
func HandleTask(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO permissiions checker
		if r.Method == http.MethodPatch {
			taskID := mux.Vars(r)["id"]
			task := checkTaskExists(s, taskID)
			if task == nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(&apierrors.TaskError{Message: "", ErrType: "Not Found"})
				fmt.Fprintln(w)
				return
			}

			decoder := json.NewDecoder(r.Body)
			updTaskJSON := &models.UpdateTaskJson{}

			if err := decoder.Decode(updTaskJSON); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&apierrors.TaskError{Message: err.Error(), ErrType: "Wrong json"})
				fmt.Fprintln(w)
				return
			}

			updateTaskFields(task, updTaskJSON)

			if err := s.EditTask(task); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(&apierrors.TaskError{Message: err.Error(), ErrType: "I am broken"})
				fmt.Fprintln(w)
				return
			}

			w.WriteHeader(http.StatusAccepted)
			return
		}

		if r.Method == http.MethodDelete {
			taskID := mux.Vars(r)["id"]
			s.DeleteTask(taskID)
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
