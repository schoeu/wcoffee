package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() == "debug" {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		c.Next()
	}
}
