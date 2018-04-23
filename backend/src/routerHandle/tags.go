package routerHandle

import (
	"../utils"
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

var (
	prefixStr = "category-"
)

func Tags(c *gin.Context, db *sql.DB) {
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
}
