package models

// User model
type User struct {
	ID       int    `bson:"id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
