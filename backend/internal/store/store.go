package store

import "github.com/dementevda/likeisaid-gg/backend/internal/api/models"

//Store interface for API storage
type Store interface {
	Open() error
	Close()
	AddUser(u *models.CreateUser) (*models.User, error)
	FindUser(login string) (*models.User, error)
}
