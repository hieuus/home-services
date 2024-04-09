package repositories

import (
	"context"
	"github.com/hieuus/home-services/internal/models"
)

type Repository interface {
	Ping() error
	Transaction(txFunc func(Repository) error) error
	UserRepository
}

type UserRepository interface {
	RecordUser(ctx context.Context, user *models.User) error
}
