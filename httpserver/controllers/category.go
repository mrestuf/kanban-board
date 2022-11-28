package controllers

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mrestuf/kanban-board/common"
	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/services"
)

type CategoryController struct {
	svc services.CategorySvc
}

func NewCategoryController(svc services.CategorySvc) *CategoryController {
	return &CategoryController{
		svc: svc,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req params.CreateCategory

	if err := ctx.ShouldBindJSON(&req); err != nil {
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
	userRole := userData.Role
	if userRole != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}
	err := validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := c.svc.CreateCategory(ctx, &req)
	WriteJsonResponse(ctx, response)
}

// func (c *CategoryController) GetCategories(ctx *gin.Context) {
// 	response:= c.svc.GetCategories(ctx)
// 	WriteJsonResponse(ctx, response)
// }