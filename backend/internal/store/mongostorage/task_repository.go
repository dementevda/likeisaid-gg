package mongostorage

import (
	"context"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddTask adds task to database
func (s *MongoStorage) AddTask(u *models.CreateTask) (*models.Task, error) {
	added, err := s.Db.Collection("tasks").InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	task := models.Task{CreateTask: u, ID: added.InsertedID.(primitive.ObjectID)}

	return &task, nil

}

// EditTask searchs task in database
func (s *MongoStorage) EditTask(task *models.Task) error {
	_, err := s.Db.Collection("tasks").ReplaceOne(context.TODO(), bson.M{"_id": task.ID}, task)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask ...
func (s *MongoStorage) DeleteTask(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := s.Db.Collection("tasks").DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

// GetTaskByID ...
func (s *MongoStorage) GetTaskByID(id string) (*models.Task, error) {
	task := &models.Task{}
	objID, _ := primitive.ObjectIDFromHex(id)

	err := s.Db.Collection("tasks").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(task)
	if err != nil {
		return nil, err
	}

	return task, err
}

// GetUserTasks ...
func (s *MongoStorage) GetUserTasks(email string) ([]*models.Task, error) {
	// TODO paginatioin and filtering
	tasks := make([]*models.Task, 0)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"wait_before", 1}})
	cur, err := s.Db.Collection("tasks").Find(context.TODO(), bson.M{"user_email": email}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		result := &models.Task{}
		err := cur.Decode(result)
		if err != nil {

			continue
		}
		tasks = append(tasks, result)

	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return tasks, err
}
