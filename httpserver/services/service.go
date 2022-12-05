package services

import (
	"context"

	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
	UpdateUser(ctx context.Context, id int, user *params.UpdateUser) *views.Response
	DeleteUser(ctx context.Context, id int) *views.Response
}

type CategorySvc interface {
	GetCategory(ctx context.Context, task *models.Task, user *models.Users) *views.Response
	CreateCategory(ctx context.Context, category *params.CreateCategory) *views.Response
	UpdateCategory(ctx context.Context, category *params.UpdateCategory, id int) *views.Response
	GetCategoryByID(ctx context.Context, id int) (*models.Categories, error)
	DeleteCategory(ctx context.Context, id int) *views.Response
}

type TaskSvc interface {
	CreateTask(ctx context.Context, task *params.CreateTask, userId int) *views.Response
	GetTasks(ctx context.Context) *views.Response
	UpdateTask(ctx context.Context, task *params.UpdateTask, taskId int) *views.Response
	UpdateTaskStatus(ctx context.Context, task *params.UpdateTaskStatus, taskId int) *views.Response
	UpdateTaskCategory(ctx context.Context, task *params.UpdateTaskCategory, taskId int) *views.Response
	DeleteTask(ctx context.Context, taskId int) *views.Response
}
