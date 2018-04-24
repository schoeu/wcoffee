package main

import (
	"./config"
	"./middlewares"
	"./routers"
	"./utils"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConf()
	r := gin.Default()

	if conf.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(middlewares.CORS())

	db := utils.OpenDb("mysql", conf.DBString)

	r.GET("/api/:type", func(c *gin.Context) {
		path := c.Param("type")
		handler := routers.RouterMap[path]

		if handler != nil {
			handler(c, db)
		} else {
			utils.ReturnError(c, "No such operation.")
		}
	})

	r.Run(":" + conf.Port)
}
