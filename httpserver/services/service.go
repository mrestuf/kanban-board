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
	UpdateUser(ctx context.Context, id int, params *params.UpdateUser) *views.Response
	DeleteUser(ctx context.Context, id int) *views.Response
}

