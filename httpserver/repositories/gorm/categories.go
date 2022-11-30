package gorm

import (
	"context"
	"time"

	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"
)

type categoriesRepo struct {
	db *gorm.DB
}

func NewCategories(db *gorm.DB) repositories.CategoriesRepo {
	return &categoriesRepo{
		db: db,
	}
}

func (r *categoriesRepo) CreateCategories(ctx context.Context, cat *models.Categories) error {
	// user := new(models.Users)
	cat.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Preload("users").Create(cat).Error
}

func (r *categoriesRepo) GetCategories(ctx context.Context) ([]models.Categories, error) {
	var categories []models.Categories

	return categories, r.db.WithContext(ctx).Find(&categories).Error
}
