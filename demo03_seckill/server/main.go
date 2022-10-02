package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"server/router"
)

var (
	port string
)

func main() {
	r := gin.New()
	r.Use(Cors()) //一定要在加路由之前加入中间件
	router.InitRouter(r)
	r.Run("localhost:" + port)
}

func init() {
	file, err := ini.Load("static/config.ini")
	if err != nil {
		panic(err)
	}
	port = file.Section("server").Key("port").String()
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("执行跨域请求处理")
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
