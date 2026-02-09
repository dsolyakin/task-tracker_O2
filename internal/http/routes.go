package http

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	r.GET("/bu", BuHandler)
}
