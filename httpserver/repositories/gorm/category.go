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

func (r *categoryRepo) FindCategoryByID(ctx context.Context, id int) (*models.Categories, error) {
	category := new(models.Categories)
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(category).Error
	return category, err
}

func (r *categoryRepo) UpdateCategory(ctx context.Context, category *models.Categories, id int) error {
	category.UpdatedAt = time.Now()
	err := r.db.WithContext(ctx).Model(category).Where("id = ?", id).Updates(*category).Error
	return err
}

func (r *categoryRepo) DeleteCategory(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Categories{}, "id = ?", id).Error
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


