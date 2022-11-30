package services

import (
	"context"
	"log"
	"net/http"

	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
	"github.com/mrestuf/kanban-board/httpserver/repositories"
	"github.com/mrestuf/kanban-board/httpserver/repositories/models"
)

type categoriesSvc struct {
	repo     repositories.CategoriesRepo
	userRepo repositories.UserRepo
}

func NewCategoriesSvc(repo repositories.CategoriesRepo, userRepo repositories.UserRepo) CategoriesSvc {
	return &categoriesSvc{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (svc *categoriesSvc) CreateCategory(ctx context.Context, cat *params.CreateCategories, UserID uint, user *params.Login) *views.Response {
	model, err := svc.userRepo.FindUserByEmail(ctx, user.Email)
	c := models.Categories{
		Type: cat.Type,
	}
	if model.Role == models.Admin {
		// return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		log.Println("only admin can create", model.Role)
		err := svc.repo.CreateCategories(ctx, &c)
		if err != nil {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}
		return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateCategories{
			Id:        c.Id,
			Type:      c.Type,
			CreatedAt: c.CreatedAt,
		})
	}

	// log.Printf("Admin Can create Category: %+v\n", model.Role)
	return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
}

func (svc *categoriesSvc) GetCategory(ctx context.Context) *views.Response {
	return nil
}
