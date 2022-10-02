package router

import (
	"github.com/gin-gonic/gin"
	. "redis/controller"
)

func InitRouter(r *gin.Engine) {
	authGroup := r.Group("/auth_code")
	{
		authGroup.POST("/getCode", AuthCodeSave)
		authGroup.POST("/judgeCode", AuthCodeJudge)
	}
}
