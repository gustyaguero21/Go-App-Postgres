package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Urlmapping(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
