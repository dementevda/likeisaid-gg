package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task model
type Task struct {
	*CreateTask `bson:",inline"`
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Notificated bool               `bson:"notifcated" json:"notifcated"`
	Done        bool               `bson:"done" json:"done"`
}

func (u *Task) String() string {
	return fmt.Sprintf("<task: %s>", u.Title)
}

// CreateTask ...
type CreateTask struct {
	*CreateTaskJson `bson:",inline"`
	UserEmail       string    `bson:"user_email" json:"user_email"`
	CreatedAt       time.Time `bson:"created_at" json:"created_at"`
}

// CreateTaskJson for creating new tasks
type CreateTaskJson struct {
	Title       string    `bson:"title" json:"title" valid:"required"`
	Description string    `bson:"description" json:"description" valid:"required"`
	WaitBefore  time.Time `bson:"wait_before" json:"wait_before" valid:"required"`
}

// UpdateTaskJson ...
type UpdateTaskJson struct {
	Title       string    `bson:"title" json:"title,omitempty"`
	Description string    `bson:"description" json:"description,omitempty"`
	WaitBefore  time.Time `bson:"wait_before" json:"wait_before,omitempty"`
}
