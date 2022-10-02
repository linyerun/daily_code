package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/service"
)

func Ready(c *gin.Context) {
	seckillService := service.NewSeckillService()
	if err := seckillService.AddProduct(); err != nil {
		c.JSON(501, gin.H{
			"msg":  err.Error(),
			"code": 501,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "成功",
		"code": http.StatusOK,
	})
}

func Go(c *gin.Context) {

}
