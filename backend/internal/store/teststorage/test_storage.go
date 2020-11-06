package teststorage

import (
	"fmt"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TestStorage implements store.Store interface
type TestStorage struct {
	Users map[string]*models.User
	Tasks map[string]*models.Task
}

func (t *TestStorage) Open() error { return nil }

func (t *TestStorage) Close() {}

func (t *TestStorage) AddUser(u *models.CreateUser) (*models.User, error) {
	if u.Login == "internalerror" {
		return nil, fmt.Errorf("internal error")
	}

	user := &models.User{
		CreateUser: u,
	}

	if _, found := t.Users[user.Email]; found {
		err := mongo.WriteException{
			WriteErrors: []mongo.WriteError{
				mongo.WriteError{
					Code: 11000,
				},
			},
		}
		return nil, err
	}

	t.Users[user.Email] = user
	return user, nil
}

func (t *TestStorage) FindUser(login string) (*models.User, error) {
	if login == "internalerror" {
		return nil, fmt.Errorf("internal error")
	}

	for _, user := range t.Users {
		if user.Login == login {
			return user, nil
		}
	}
	return nil, mongo.ErrNoDocuments
}

func (t *TestStorage) GetUserByEmail(email string) (*models.User, error) {
	if user, found := t.Users[email]; found {
		return user, nil
	}
	return nil, fmt.Errorf("%s", "Not Found")
}

func (t *TestStorage) AddTask(create *models.CreateTask) (*models.Task, error) {
	if create.Title == "internalerror" {
		return nil, fmt.Errorf("internal error")
	}

	task := &models.Task{CreateTask: create}
	task.ID = primitive.NewObjectID()
	t.Tasks[task.ID.String()] = task
	return task, nil

}

func (t *TestStorage) GetUserTasks(email string, filters *store.TaskFilters) ([]*models.Task, error) {
	if email == "internalerror" {
		return nil, fmt.Errorf("Internal error")
	}

	var tasks []*models.Task

	for _, task := range t.Tasks {
		if task.UserEmail == email {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (t *TestStorage) GetTaskByID(id string) (*models.Task, error) {
	if task, found := t.Tasks[id]; found {
		return task, nil
	}
	return nil, fmt.Errorf("%s", "Not Found")
}
func (t *TestStorage) EditTask(task *models.Task) error {
	if task.Title == "internalerror" {
		return fmt.Errorf("Internal error")
	}
	return nil
}

func (t *TestStorage) DeleteTask(id string) error {
	delete(t.Tasks, id)
	return nil
}
