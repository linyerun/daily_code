package main

import (
	"github.com/gin-gonic/gin"
	. "redis/router"
)

func main() {
	e := gin.New()
	InitRouter(e)
	if err := e.Run("localhost:9090"); err != nil {
		panic(err)
	}
}
