package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store for API
type Store struct {
	config *Config
	client *mongo.Client
	db     *mongo.Database
}

//New returns Store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open connection to store
func (s *Store) Open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(s.config.DatabaseURL))
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
	s.db = client.Database("likeisaid")

	return nil
}

// Close connection to store
func (s *Store) Close() {
	s.client.Disconnect(context.TODO())
}
