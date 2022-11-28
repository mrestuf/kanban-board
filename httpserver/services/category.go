package services

import (
	"context"
	"net/http"

	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
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
		Id:             c.Id,
		Type:           c.Type,
		CreatedAt:      c.CreatedAt,
	})
	
}

// func (s *categorySvc) GetCategories(ctx context.Context) *views.Response {
// 	c, err := s.repo.GetCategories(ctx)
// 	if err != nil {
// 		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
// 	}
// }