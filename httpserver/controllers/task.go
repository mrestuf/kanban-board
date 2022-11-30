package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mrestuf/kanban-board/common"
	"github.com/mrestuf/kanban-board/httpserver/controllers/params"
	"github.com/mrestuf/kanban-board/httpserver/services"
	"net/http"
	"strconv"
)

type TaskController struct {
	svc services.TaskSvc
}

func NewTaskController(svc services.TaskSvc) *TaskController {
	return &TaskController{
		svc: svc,
	}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var req params.CreateTask
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
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

	response := c.svc.CreateTask(ctx, &req, userData.Id)
	WriteJsonResponse(ctx, response)
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	response := c.svc.GetTasks(ctx)
	WriteJsonResponse(ctx, response)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdateTask
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.UpdateTask(ctx, &req, id)
	WriteJsonResponse(ctx, response)
}

func (c *TaskController) UpdateTaskStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdateTaskStatus
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.UpdateTaskStatus(ctx, &req, id)
	WriteJsonResponse(ctx, response)
}

func (c *TaskController) UpdateTaskCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdateTaskCategory
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.UpdateTaskCategory(ctx, &req, id)
	WriteJsonResponse(ctx, response)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.DeleteTask(ctx, id)
	WriteJsonResponse(ctx, response)
}
