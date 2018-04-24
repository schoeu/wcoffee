package routerHandle

import (
	"../utils"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type OptScruct struct {
	Name string `json:"name"`
	List []string `json:"list"`
	Default string `json:"default"`
}

func Options(c *gin.Context, db *sql.DB) {
	options := OptScruct{}
	catevalId := c.Query("id")
	if catevalId != "" {
		rows, err := db.Query("SELECT name, price, attrs_name, is_default FROM tb_attrs_value where cateval_id = ? order by attrs_id", catevalId)
		utils.ErrHandle(err)

		var name, attrName sql.NullString
		var price sql.NullInt64
		for rows.Next() {
			err := rows.Scan(&name, &price, &attrName)
			utils.ErrHandle(err)

			options.Name = attrName
			options.Default
		}
		err = rows.Err()
		utils.ErrHandle(err)
		defer rows.Close()
	}
	c.JSON(200, gin.H{
		"status": 0,
		"data":   "",
	})
}
