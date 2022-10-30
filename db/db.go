package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/upms")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}
