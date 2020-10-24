package mongostorage

import (
	"context"

	"github.com/dementevda/likeisaid-gg/backend/cmd/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStorage for API
type MongoStorage struct {
	config *store.Config
	client *mongo.Client
	Db     *mongo.Database
}

//New ...
func New(config *store.Config) *MongoStorage {
	return &MongoStorage{
		config: config,
	}
}

// Open connection to storage
func (s *MongoStorage) Open() error {
	opts := options.Client().ApplyURI(s.config.DatabaseURL).SetAuth(options.Credential{AuthSource: s.config.DatabaseName, Username: s.config.DatabaseUser, Password: s.config.DatabasePasswd})
	client, err := mongo.NewClient(opts)
	if err != nil {
		return err
	}

	if err := client.Connect(context.TODO()); err != nil {
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	s.client = client
	s.Db = client.Database(s.config.DatabaseName)

	return nil
}

// Close connection to storage
func (s *MongoStorage) Close() {
	s.client.Disconnect(context.TODO())
}
