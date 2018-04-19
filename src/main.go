package main

import (
	"./config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConf()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + conf.Port)
}
