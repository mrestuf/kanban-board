package gorm

import (
	"context"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"
	"time"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) repositories.TaskRepo {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) CreateTask(ctx context.Context, task *models.Task) error {
	task.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepo) FindAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	if err := r.db.WithContext(ctx).Joins("User").Find(&tasks).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *taskRepo) FindTaskByID(ctx context.Context, id int) (*models.Task, error) {
	task := new(models.Task)
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(task).Error
	return task, err
}

func (r *taskRepo) UpdateTask(ctx context.Context, task *models.Task) error {
	task.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *taskRepo) DeleteTask(ctx context.Context, taskId int) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, "id = ?", taskId).Error
}
