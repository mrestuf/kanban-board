package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrestuf/kanban-board/httpserver/controllers/views"
)

func WriteJsonResponse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}