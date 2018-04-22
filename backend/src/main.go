package main

import (
	"./config"
	"./middlewares"
	"./utils"
	"github.com/gin-gonic/gin"
)

type cateInfo struct {
	Name   string `json:"name"`
	Anchor string `json:"anchor"`
}

type tagsStruct struct {
	Items []cateInfo `json:"items"`
}

type ItemStruct struct {
	Price int    `json:"price"`
	Img   string `json:"img"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
}
type listStruct struct {
	Anchor   string       `json:"anchor"`
	List     []ItemStruct `json:"list"`
	Typename string       `json:"typename"`
}

func main() {
	conf := config.GetConf()
	r := gin.Default()

	if conf.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(middlewares.CORS())

	db := utils.OpenDb("mysql", conf.DBString)

	prefixStr := "category-"

	r.GET("/api/tags", func(c *gin.Context) {
		var ts tagsStruct
		rows, err := db.Query("select category_id, name from tb_category")
		utils.ErrHandle(err)

		var id, name string
		for rows.Next() {
			err := rows.Scan(&id, &name)
			utils.ErrHandle(err)
			ci := cateInfo{}
			ci.Name = name
			ci.Anchor = prefixStr + id
			ts.Items = append(ts.Items, ci)
		}
		err = rows.Err()
		utils.ErrHandle(err)
		defer rows.Close()

		c.JSON(200, gin.H{
			"status": 0,
			"data":   ts,
		})
	})

	r.GET("/api/list", func(c *gin.Context) {
		var ls listStruct
		var lsCtt []listStruct
		rows, err := db.Query(`SELECT 
			tb_category.category_id,
			tb_category.name, 
			tb_category_values.catevalue_id, 
			tb_category_values.name, 
			tb_category_values.img, 
			tb_category_values.price, 
			tb_category_values.alias 
			FROM 
			tb_category, 
			tb_category_values
			WHERE 
			tb_category.category_id = tb_category_values.category_id
			ORDER BY
			category_id asc`)

		utils.ErrHandle(err)

		var cateId, cateName, cateValName, img, alias string
		var cateValId, price int
		for rows.Next() {
			err := rows.Scan(&cateId, &cateName, &cateValId, &cateValName, &img, &price, &alias)
			utils.ErrHandle(err)
			ls.Anchor = prefixStr + cateId
			ls.Typename = cateName
			ls.List = append(ls.List, ItemStruct{
				Id:    cateValId,
				Img:   img,
				Price: price,
				Name:  cateValName,
				Alias: alias,
			})
			lsCtt = append(lsCtt, ls)
		}
		err = rows.Err()
		utils.ErrHandle(err)
		defer rows.Close()

		c.JSON(200, gin.H{
			"status": 0,
			"data":   lsCtt,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 0,
			"data":   "I`m ready for that, you know.",
		})
	})

	r.Run(":" + conf.Port)
}
