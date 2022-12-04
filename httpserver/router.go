package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrestuf/kanban-board/common"
	"github.com/mrestuf/kanban-board/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	user     *controllers.UserController
	category *controllers.CategoryController
	task     *controllers.TaskController
}

func NewRouter(r *gin.Engine, user *controllers.UserController, category *controllers.CategoryController, task *controllers.TaskController) *router {
	return &router{
		router:   r,
		user:     user,
		category: category,
		task:     task,
	}
}

func (r *router) Start(port string) {
	//users
	r.router.POST("/v1/users/register", r.user.Register)
	r.router.POST("/v1/users/login", r.user.Login)
	r.router.PUT("/v1/users/update-account", r.verifyToken, r.user.UpdateUser)
	r.router.DELETE("/v1/users/delete-account", r.verifyToken, r.user.DeleteUser)

	//categories
	r.router.POST("/v1/categories", r.verifyToken, r.category.CreateCategory)
	r.router.GET("/v1/categrories", r.verifyToken, r.category.GetCategories)
	r.router.PATCH("/v1/categories/:categoryId", r.verifyToken, r.category.UpdateCategory)
	r.router.DELETE("/v1/categories/:categoryId", r.verifyToken, r.category.DeleteCategory)

	//task
	r.router.POST("/v1/tasks", r.verifyToken, r.task.CreateTask)
	r.router.GET("/v1/tasks", r.verifyToken, r.task.GetTasks)
	r.router.PUT("/v1/tasks/:taskId", r.verifyToken, r.task.UpdateTask)
	r.router.PATCH("/v1/tasks/update-status/:taskId", r.verifyToken, r.task.UpdateTaskStatus)
	r.router.PATCH("/v1/tasks/update-category/:taskId", r.verifyToken, r.task.UpdateTaskCategory)
	r.router.DELETE("/v1/tasks/:taskId", r.verifyToken, r.task.DeleteTask)

	r.router.Run(port)
}

func (r *router) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
}
