package router

import (
	"github.com/gin-gonic/gin"
	. "server/controller"
)

func InitRouter(r *gin.Engine) {
	seckillGroup := r.Group("/seckill/product")
	{
		seckillGroup.POST("/ready", Ready)
		seckillGroup.POST("/go/:userId", Go)
		seckillGroup.POST("/test/go", TestGo)
	}
}
