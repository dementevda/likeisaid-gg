package store

import "github.com/dementevda/likeisaid-gg/backend/internal/api/models"

//Store interface for API storage
type Store interface {
	Open() error
	Close()
	AddUser(u *models.CreateUser) (*models.User, error)
	FindUser(login string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	AddTask(*models.CreateTask) (*models.Task, error)
	GetUserTasks(string, *TaskFilters) ([]*models.Task, error)
	GetTaskByID(string) (*models.Task, error)
	EditTask(*models.Task) error
	DeleteTask(string) error
}

type TaskFilters struct {
	Done  bool
	Limit int64
	Page  int64
}

func (f *TaskFilters) Defaults() {
	f.Done = false
	f.Limit = 10
	f.Page = 1
}
