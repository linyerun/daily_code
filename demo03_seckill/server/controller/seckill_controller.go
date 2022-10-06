package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"server/service"
	"strconv"
	"time"
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
	//为了测试方便
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(501, gin.H{
			"msg":  "userId这个参数有误",
			"code": 501,
		})
		log.Println(userId, ": ", "userId这个参数有误")
		return
	}
	msg, err := service.NewSeckillService().SeckillProduct(userId)
	if err != nil {
		c.JSON(502, gin.H{
			"code": 502,
			"msg":  err.Error(),
		})
		log.Println(userId, ": ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
	log.Println(userId, ": ", msg)
}

// TestGo 为了方便测试，改造一下Go
func TestGo(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	userId := int64(rand.Intn(10000))
	msg, err := service.NewSeckillService().SeckillProduct(userId)
	if err != nil {
		c.JSON(502, gin.H{
			"code": 502,
			"msg":  err.Error(),
		})
		log.Println(userId, ": ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
	log.Println(userId, ": ", msg)
}
