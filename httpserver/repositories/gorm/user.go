package gorm

import (
	"context"

	"strings"
	"time"

	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"

	"time"
	// "strings"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.Users) error {
	user.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) FindUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := new(models.Users)
	err := r.db.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error
	return user, err
}

func (r *userRepo) FindUserByID(ctx context.Context, id uint) (*models.Users, error) {
	user := new(models.Users)
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(user).Error
	return user, err
}

func (r *userRepo) UpdateUser(ctx context.Context, user *models.Users) error {
	user.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Model(user).Updates(*user).Error
}

func (r *userRepo) DeleteUser(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Users{}, "id = ?", id).Error
}
