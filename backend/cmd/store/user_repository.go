package store

import (
	"context"
	"fmt"

	"github.com/dementevda/likeisaid-gg/backend/cmd/api/models"
)

func (s *MongoStorage) AddUser(u *models.User) (*models.User, error) {
	add, err := s.Db.Collection("users").InsertOne(context.TODO(), u)
	fmt.Println(add)
	if err != nil {
		return nil, err
	}
	return u, nil

}
