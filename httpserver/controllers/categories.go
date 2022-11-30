package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mrestuf/kanban-board/common"
	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/services"
)

type CategoriesController struct {
	svc     services.CategoriesSvc
	userSvc services.UserSvc
}

func NewCategoriesController(svc services.CategoriesSvc, userSvc services.UserSvc) *CategoriesController {
	return &CategoriesController{
		svc:     svc,
		userSvc: userSvc,
	}
}

func (control *CategoriesController) CreateCategories(ctx *gin.Context) {
	var request params.CreateCategories
	var user params.Login
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, exist := ctx.Get("userData")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}
	userData := claims.(*common.CustomClaims)
	userId := uint(userData.Id)
	err := validator.New().Struct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := control.svc.CreateCategory(ctx, &request, userId, &user)
	WriteJsonResponse(ctx, response)
}
