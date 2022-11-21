package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type router struct {
	router *gin.Engine
}

func NewRouter(r *gin.Engine) *router {
	return &router{
		router: r,
	}
}

func (r *router) Start(port string) {

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
