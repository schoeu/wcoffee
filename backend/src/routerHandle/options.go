package routerHandle

import (
	"../utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type OptItems struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type OptScruct struct {
	Name    string     `json:"name"`
	List    []OptItems `json:"list"`
	Default string     `json:"default"`
}

func Options(c *gin.Context, db *sql.DB) {
	catevalId := c.Query("id")
	var optCtt []OptScruct
	if catevalId != "" {
		rows, err := db.Query("SELECT name, price, attrs_name, is_default FROM tb_attrs_value where cateval_id = ? order by attrs_id", catevalId)
		utils.ErrHandle(err)

		var prevType string
		var name, attrName sql.NullString
		var price, isDefault sql.NullInt64
		options := OptScruct{}
		for rows.Next() {
			err := rows.Scan(&name, &price, &attrName, &isDefault)
			utils.ErrHandle(err)

			if attrName.String != prevType {
				if prevType != "" {
					optCtt = append(optCtt, options)
				}
				options = OptScruct{}
				prevType = attrName.String
			}

			if attrName.String == prevType {
				if int(isDefault.Int64) == 1 {
					options.Default = name.String
				}
				options.Name = attrName.String
				options.List = append(options.List, OptItems{
					Name:  name.String,
					Price: int(price.Int64),
				})
			}
		}

		if len(options.List) > 0 {
			optCtt = append(optCtt, options)
		}

		err = rows.Err()
		utils.ErrHandle(err)
		defer rows.Close()
	}

	c.JSON(200, gin.H{
		"status": 0,
		"data":   optCtt,
	})
}
