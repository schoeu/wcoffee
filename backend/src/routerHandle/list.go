package routerHandle

import (
	"../utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

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

func List(c *gin.Context, db *sql.DB) {
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
}
