package main

import (
	"./config"
	"./utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type cateInfo struct {
	Name   string `json:"name"`
	Anchor string `json:"anchor"`
}

type tagsStruct struct {
	Status int `json:"status"`
	Data   struct {
		Items []cateInfo `json:"items"`
	} `json:"data"`
}

func main() {
	conf := config.GetConf()
	r := gin.Default()
	//r.Use(middlewares.JsonP())

	db := utils.OpenDb("mysql", conf.DBString)

	prefixStr := "category-"

	r.GET("/api/tags", func(c *gin.Context) {
		var ts tagsStruct
		rows, err := db.Query("select category_id, name from tb_category")
		utils.ErrHandle(err)

		cbFnName := c.Query("callback")
		if cbFnName == "" {
			cbFnName = "callback"
		}

		var id, name string
		for rows.Next() {
			err := rows.Scan(&id, &name)
			utils.ErrHandle(err)
			ci := cateInfo{}
			ci.Name = name
			ci.Anchor = prefixStr + id
			ts.Data.Items = append(ts.Data.Items, ci)
		}
		err = rows.Err()
		utils.ErrHandle(err)
		defer rows.Close()

		ts.Status = 0
		b, err := json.Marshal(ts)
		utils.ErrHandle(err)
		fmt.Fprintf(c.Writer, "%s(%s);", cbFnName, b)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 0,
			"data":   "I`m ready for that, you know.",
		})
	})

	r.Run(":" + conf.Port)
}
