package mongostorage

import (
	"context"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddUser adds user to database
func (s *MongoStorage) AddUser(u *models.CreateUser) (*models.User, error) {
	added, err := s.Db.Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	user := models.User{CreateUser: u, ID: added.InsertedID.(primitive.ObjectID)}

	return &user, nil

}

// FindUser search user in database
func (s *MongoStorage) FindUser(login string) (*models.User, error) {
	user := &models.User{}

	err := s.Db.Collection("users").FindOne(context.TODO(), bson.M{"login": login}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, err
}
