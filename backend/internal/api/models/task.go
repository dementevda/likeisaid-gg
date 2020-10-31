package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task model
type Task struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	UserEmail   string              `bson:"user_email"`
	Title       string              `bson:"title"`
	Defendant   string              `bson:"defendant"`
	Description string              `bson:"description,omitempty"`
	CreatedAt   primitive.Timestamp `bson:"created_at"`
	WaitBefore  time.Time           `bson:"wait_before"`
}

func (u *Task) String() string {
	return fmt.Sprintf("<task: %s>", u.Title)
}
