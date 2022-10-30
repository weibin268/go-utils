package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/weibin268/go-utils/io"
	"log"
	"os"
)

var Db *sqlx.DB

func init() {
	// 读取配置文件
	config := viper.New()
	config.SetConfigName("db.ini")
	config.SetConfigType("ini")
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	driverName := config.GetString("db.driverName") // 读取配置
	dataSourceName := config.GetString("db.dataSourceName")
	// 创建数据库链接
	database, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func ExecSqlFromFile() {
	fileName := "db.sql"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	strSql := io.ReadText(fileName)
	result, err := Db.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}
