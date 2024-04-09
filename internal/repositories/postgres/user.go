package postgres

import (
	"context"
	"github.com/hieuus/home-services/internal/models"
)

func (q *Queries) RecordUser(ctx context.Context, user *models.User) error {
	return q.db.WithContext(ctx).Model(&models.User{}).Create(user).Error
}
