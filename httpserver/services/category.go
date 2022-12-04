package services

import (
	"context"
	"net/http"

	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
	"gorm.io/gorm"
)

type categorySvc struct {
	repo repositories.CategoryRepo
	// user repositories.UserRepo
}

func NewCategorySvc(repo repositories.CategoryRepo) CategorySvc {
	return &categorySvc{
		repo: repo,
	}
}

func (s *categorySvc) CreateCategory(ctx context.Context, category *params.CreateCategory) *views.Response {
	//req
	c := models.Categories{
		Type: category.Type,
	}

	err := s.repo.CreateCategory(ctx, &c)

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateCategory{
		Id:        c.Id,
		Type:      c.Type,
		CreatedAt: c.CreatedAt,
	})
}

func (s *categorySvc) GetCategoryByID(ctx context.Context, id int) (*models.Categories, error) {
	c, err := s.repo.FindCategoryByID(ctx, id)
	if err != nil {
		return c, err
	}
	if c.Id == 0 {
		return c, nil
	}

	return c, nil
}

func (s *categorySvc) UpdateCategory(ctx context.Context, category *params.UpdateCategory, id int) *views.Response {
	c, err := s.repo.FindCategoryByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	c.Type = category.Type

	err = s.repo.UpdateCategory(ctx, c, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateCategory{
		Id:        c.Id,
		Type:      c.Type,
		UpdatedAt: c.UpdatedAt,
	})
}

func (s *categorySvc) DeleteCategory(ctx context.Context, id int) *views.Response {
	_, err := s.repo.FindCategoryByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = s.repo.DeleteCategory(ctx, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_CATEGORY_SUCCESSFULLY_DELETED, nil)
}

func (s *categorySvc) GetCategory(ctx context.Context) *views.Response {
	c, err := s.repo.GetCategories(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	categories := make([]views.GetCategories, 0)
	for _, cat := range c {
		categories = append(categories, views.GetCategories{
			Id:        cat.Id,
			Type:      cat.Type,
			UpdatedAt: cat.UpdatedAt,
			CreatedAt: cat.CreatedAt,
			Tasks: views.TaskGetCategories{
				Id:          cat.TaskId,
				Title:       cat.Task.Title,
				Description: cat.Task.Description,
				UserId:      cat.UserId,
				CategoryId:  cat.CategoryId,
				UpdatedAt:   cat.Task.UpdatedAt,
				CreatedAt:   cat.Task.CreatedAt,
			},
		})
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, categories)
}
