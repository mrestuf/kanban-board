package gorm

import (
	"context"
	// "strings"
	"time"
	"log"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"

)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) repositories.CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) CreateCategory(ctx context.Context, category *models.Categories) error {
	category.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepo) GetCategories(ctx context.Context) ([]models.Categories, error) {
	var categories []models.Categories

	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		log.Println(err)
		return categories, err
	}
	return categories, nil
}


