package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func (u *User) String() string {
	return fmt.Sprintf("%s, %s, %s", u.ID, u.Email, u.Password)
}
