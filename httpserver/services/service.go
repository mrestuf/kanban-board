package services

import (
	"context"

	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	// "github.com/mrestuf/kanban-board/httpserver/repositories/models"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
	UpdateUser(ctx context.Context, id int, user *params.UpdateUser) *views.Response
	DeleteUser(ctx context.Context, id int) *views.Response
}

type CategoriesSvc interface {
	CreateCategory(ctx context.Context, cat *params.CreateCategories, UserID uint, user *params.Login) *views.Response
	GetCategory(ctx context.Context) *views.Response
}
