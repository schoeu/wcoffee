package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


// 统一错误处理
func ErrHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 创建数据库链接
func OpenDb(dbTyepe string, dbStr string) *sql.DB {
	if dbTyepe == "" {
		dbTyepe = "mysql"
	}
	db, err := sql.Open(dbTyepe, dbStr)
	ErrHandle(err)

	err = db.Ping()
	ErrHandle(err)
	return db
}
