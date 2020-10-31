package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	*CreateUser `bson:",inline"`
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Tasks       []primitive.ObjectID `bson:"tasks" json:"tasks"`
}

func (u *User) String() string {
	return fmt.Sprintf("<user: %s>", u.Email)
}

// CreateUser for creating new users
type CreateUser struct {
	Email string `bson:"email" json:"email"  valid:"required,email"`
	Login string `bson:"login" json:"login" valid:"required"`
	// AuthEngine string `bson:"auth_engine" json:"auth_engine" valid:"required,in(google|yandex|github|self)"`
}
