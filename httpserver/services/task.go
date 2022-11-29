package services

import (
	"context"
	"errors"
	"github.com/mrestuf/kanban-board/common"
	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"
	"net/http"
)

type taskSvc struct {
	repo         repositories.TaskRepo
	categoryRepo repositories.CategoryRepo
}

func NewTaskSvc(repo repositories.TaskRepo, categoryRepo repositories.CategoryRepo) TaskSvc {
	return &taskSvc{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (s *taskSvc) CreateTask(ctx context.Context, task *params.CreateTask, userId int) *views.Response {
	if _, err := s.categoryRepo.FindCategoryByID(ctx, task.CategoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model := models.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryId:  task.CategoryId,
		Status:      false,
		UserId:      userId,
	}

	err := s.repo.CreateTask(ctx, &model)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateTask{
		Id:          model.Id,
		Title:       model.Title,
		Status:      model.Status,
		Description: model.Description,
		UserId:      model.UserId,
		CategoryId:  model.CategoryId,
		CreatedAt:   model.CreatedAt,
	})
}

func (s *taskSvc) GetTasks(ctx context.Context) *views.Response {
	tasks, err := s.repo.FindAllTasks(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	resp := make([]views.FindAllTasks, len(tasks))
	for i, task := range tasks {
		r := views.FindAllTasks{
			Id:          task.Id,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserId:      task.UserId,
			CategoryId:  task.CategoryId,
			CreatedAt:   task.CreatedAt,
		}
		r.User.Id = task.UserId
		r.User.FullName = task.User.FullName
		r.User.Email = task.User.Email
		resp[i] = r
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, resp)
}

func (s *taskSvc) UpdateTask(ctx context.Context, params *params.UpdateTask, taskId int) *views.Response {
	task, err := s.repo.FindTaskByID(ctx, taskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if task.UserId != ctx.Value("userData").(*common.CustomClaims).Id {
		return views.ErrorResponse(http.StatusUnauthorized, views.M_UNAUTHORIZED_ACTION, errors.New("unauthorized to update other user task"))
	}

	task.Title = params.Title
	task.Description = params.Description

	err = s.repo.UpdateTask(ctx, task)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateTask{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserId:      task.UserId,
		CategoryId:  task.CategoryId,
		UpdatedAt:   task.UpdatedAt,
	})
}

func (s *taskSvc) UpdateTaskStatus(ctx context.Context, params *params.UpdateTaskStatus, taskId int) *views.Response {
	task, err := s.repo.FindTaskByID(ctx, taskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if task.UserId != ctx.Value("userData").(*common.CustomClaims).Id {
		return views.ErrorResponse(http.StatusUnauthorized, views.M_UNAUTHORIZED_ACTION, errors.New("unauthorized to update other user task"))
	}

	task.Status = *params.Status

	err = s.repo.UpdateTask(ctx, task)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateTask{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserId:      task.UserId,
		CategoryId:  task.CategoryId,
		UpdatedAt:   task.UpdatedAt,
	})
}

func (s *taskSvc) UpdateTaskCategory(ctx context.Context, params *params.UpdateTaskCategory, taskId int) *views.Response {
	task, err := s.repo.FindTaskByID(ctx, taskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if task.UserId != ctx.Value("userData").(*common.CustomClaims).Id {
		return views.ErrorResponse(http.StatusUnauthorized, views.M_UNAUTHORIZED_ACTION, errors.New("unauthorized to update other user task"))
	}

	if _, err := s.categoryRepo.FindCategoryByID(ctx, params.CategoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	task.CategoryId = params.CategoryId

	err = s.repo.UpdateTask(ctx, task)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateTask{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserId:      task.UserId,
		CategoryId:  task.CategoryId,
		UpdatedAt:   task.UpdatedAt,
	})
}

func (s *taskSvc) DeleteTask(ctx context.Context, taskId int) *views.Response {
	task, err := s.repo.FindTaskByID(ctx, taskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if task.UserId != ctx.Value("userData").(*common.CustomClaims).Id {
		return views.ErrorResponse(http.StatusUnauthorized, views.M_UNAUTHORIZED_ACTION, errors.New("unauthorized to delete other user task"))
	}

	err = s.repo.DeleteTask(ctx, taskId)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_TASK_SUCCESSFULLY_DELETED, nil)
}
