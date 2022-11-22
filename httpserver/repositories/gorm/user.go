package gorm

import (
	"context"

	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"
	"time"
	"strings"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := new(models.User)
	err := r.db.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error
	return user, err
}