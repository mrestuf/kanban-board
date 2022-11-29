package repositories

import (
	"context"

	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindUserByEmail(ctx context.Context, email string) (*models.Users, error)
	FindUserByID(ctx context.Context, id int) (*models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) error
	DeleteUser(ctx context.Context, id int) error
}

type CategoryRepo interface {
	GetCategories(ctx context.Context) ([]models.Categories, error)
	CreateCategory(ctx context.Context, category *models.Categories) error
	FindCategoryByID(ctx context.Context, id int) (*models.Categories, error)
	UpdateCategory(ctx context.Context, category *models.Categories, id int) error
	DeleteCategory(ctx context.Context, id int) error
}
