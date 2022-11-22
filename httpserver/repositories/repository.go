package repositories

import (
	"context"

	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
}