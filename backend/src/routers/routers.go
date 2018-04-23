package routers

import (
	"../routerHandle"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var RouterMap = map[string]func(*gin.Context, *sql.DB){
	"tags":    routerHandle.Tags,
	"options": routerHandle.Options,
	"list":    routerHandle.List,
}
