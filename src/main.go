package main

import (
	"github.com/gin-gonic/gin"
	"./config"
)

func main() {
	conf := config.GetConf()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "I`m ready for that, you know.",
		})
	})

	r.Run(":" + conf.Port)
}