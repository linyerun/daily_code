package router

import (
	"github.com/gin-gonic/gin"
	. "server/controller"
)

func InitRouter(r *gin.Engine) {
	authGroup := r.Group("/seckill/product")
	{
		authGroup.POST("/ready", Ready)
		authGroup.POST("/go", Go)
	}
}
