package mongostorage

import (
	"context"

	"github.com/dementevda/likeisaid-gg/backend/internal/api/models"
	"github.com/dementevda/likeisaid-gg/backend/internal/store"
	pag "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddTask adds task to database
func (s *MongoStorage) AddTask(t *models.CreateTask) (*models.Task, error) {
	task := &models.Task{CreateTask: t}
	added, err := s.Db.Collection("tasks").InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}
	task.ID = added.InsertedID.(primitive.ObjectID)
	return task, nil

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
func (s *MongoStorage) GetUserTasks(email string, filters *store.TaskFilters) ([]*models.Task, error) {
	filter := bson.M{"user_email": email, "done": filters.Done}
	collection := s.Db.Collection("tasks")
	paginatedData, err := pag.New(collection).Limit(filters.Limit).Page(filters.Page).Sort("wait_before", 1).Filter(filter).Find()

	if err != nil {
		return nil, err
	}

	var taskSlice []*models.Task
	for _, raw := range paginatedData.Data {
		task := &models.Task{}
		marshallErr := bson.Unmarshal(raw, task)
		if marshallErr == nil {
			taskSlice = append(taskSlice, task)

		}
	}

	return taskSlice, err
}
