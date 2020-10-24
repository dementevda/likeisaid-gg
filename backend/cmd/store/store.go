package store

import "github.com/dementevda/likeisaid-gg/backend/cmd/api/models"

//Store interface for API storage
type Store interface {
	Open() error
	Close()
	AddUser(u *models.User) (*models.User, error)
	FindUser(u *models.User) (*models.User, error)
}
