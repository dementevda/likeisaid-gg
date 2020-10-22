package models

import "go.mongodb.org/mongo-driver/mongo"

// User model
type User struct {
	db *mongo.Collection

	ID       int
	Email    string
	Password string
}
