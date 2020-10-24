package mongostorage

import (
	"context"

	"github.com/dementevda/likeisaid-gg/backend/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddUser adds user to database
func (s *MongoStorage) AddUser(u *models.User) (*models.User, error) {
	u.Email = "qwe@qwe"
	added, err := s.Db.Collection("users").InsertOne(context.TODO(), u)
	u.ID = added.InsertedID.(primitive.ObjectID)
	if err != nil {
		return nil, err
	}
	return u, nil

}

// FindUser search user in database
func (s *MongoStorage) FindUser(u *models.User) (*models.User, error) {
	user := &models.User{}
	err := s.Db.Collection("users").FindOne(context.TODO(), bson.M{"email": u.Email}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}
