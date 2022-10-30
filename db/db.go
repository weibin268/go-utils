package db

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/weibin268/go-utils/io"
	"log"
	"sort"
)

var Db *sqlx.DB

func InitDb() {
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

func CloseDb() {
	Db.Close()
}

func QueryMaps(sqlStr string, args ...interface{}) ([]map[string]string, error) {
	rows, err := Db.Query(sqlStr, args...)
	if err != nil {
		return nil, err
	}
	//函数结束释放链接
	defer rows.Close()
	//读出查询出的列字段名
	cols, _ := rows.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面,狸猫换太子
	for i := range values {
		scans[i] = &values[i]
	}
	results := make([]map[string]string, 0, 10)
	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]string, 10)
		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	//返回数据
	return results, nil
}

func ExecSqlFromFile() {
	// 命令行参数
	var sqlFile string
	var sqlType string
	flag.StringVar(&sqlFile, "f", "db.sql", "sql file")
	flag.StringVar(&sqlType, "t", "", "sql type: query、update")
	flag.Parse()

	strSql := io.ReadText(sqlFile)
	if sqlType == "query" {
		// 查询
		rows, err := QueryMaps(strSql)
		if err != nil {
			log.Fatal(err)
		}
		var colums []string
		for _, r := range rows {
			if colums == nil {
				for k, _ := range r {
					colums = append(colums, k)
				}
				sort.Strings(colums)
				for _, c := range colums {
					fmt.Printf("%v|", c)
				}
				fmt.Printf("\n")
			}
			for _, k := range colums {
				s := r[k]
				fmt.Printf("%v|", s)
			}
			fmt.Printf("\n")
		}
	} else {
		// 更新
		result, err := Db.Exec(strSql)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.RowsAffected())
	}
}
