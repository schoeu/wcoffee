package main

import (
	"./config"
	"./middlewares"
	"./utils"
	"database/sql"
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
	Desc  string `json:"desc"`
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

	r.GET("/api/options", func(c *gin.Context) {
		var ts tagsStruct
		catevalId := c.Query("id")
		if catevalId != "" {
			var name, attrName sql.NullString
			var price sql.NullInt64
			rows, err := db.Query("SELECT name, price, attrs_name FROM tb_attrs_value where cateval_id = ? order by attrs_id", catevalId)
			utils.ErrHandle(err)
	
			var id, name string
			for rows.Next() {
				err := rows.Scan(&name, &price, &attrName)
				utils.ErrHandle(err)
				ci := cateInfo{}
				ci.Name = name
				ci.Anchor = prefixStr + id
				ts.Items = append(ts.Items, ci)
			}
			err = rows.Err()
			utils.ErrHandle(err)
			defer rows.Close()
		}
		c.JSON(200, gin.H{
			"status": 0,
			"data":   ts,
		})
	})

	r.GET("/api/list", func(c *gin.Context) {
		var lsCtt []listStruct
		rows, err := db.Query(`SELECT 
			tb_category.category_id,
			tb_category.name, 
			tb_category_values.catevalue_id, 
			tb_category_values.name, 
			tb_category_values.img, 
			tb_category_values.price, 
			tb_category_values.alias,
			tb_category_values.desc 
			FROM 
			tb_category, 
			tb_category_values
			WHERE 
			tb_category.category_id = tb_category_values.category_id
			ORDER BY
			category_id asc`)

		utils.ErrHandle(err)

		var cateId, cateName, cateValName, img, alias, desc sql.NullString
		var cateValId, price sql.NullInt64
		var preType = "1"
		var ls listStruct
		for rows.Next() {
			err := rows.Scan(&cateId, &cateName, &cateValId, &cateValName, &img, &price, &alias, &desc)
			utils.ErrHandle(err)

			if cateId.String != preType {
				lsCtt = append(lsCtt, ls)
				preType = cateId.String
				ls.List = []ItemStruct{}
			} else {
				ls.Anchor = prefixStr + cateId.String
				ls.Typename = cateName.String
				ls.List = append(ls.List, ItemStruct{
					Id:    int(cateValId.Int64),
					Img:   img.String,
					Price: int(price.Int64),
					Name:  cateValName.String,
					Alias: alias.String,
					Desc:  desc.String,
				})
			}

		}
		if len(ls.List) > 0 {
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
